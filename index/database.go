package index

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DbClient struct {
	Pool *pgxpool.Pool
}

func afterConnectRegisterTypes(ctx context.Context, conn *pgx.Conn) error {
	data_type_names := []string{
		"tonaddr",
		"_tonaddr",
		"tonhash",
		"_tonhash",
	}
	for _, type_name := range data_type_names {
		data_type, _ := conn.LoadType(ctx, type_name)
		conn.TypeMap().RegisterType(data_type)
	}
	return nil
}

func NewDbClient(dsn string, maxconns int, minconns int) (*DbClient, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	
	// Set reasonable defaults for external Postgres connection
	if maxconns > 0 {
		config.MaxConns = int32(maxconns)
	} else {
		config.MaxConns = 20 // Lower default for external connections
	}
	
	if minconns > 0 {
		config.MinConns = int32(minconns)
	} else {
		config.MinConns = 2 // Ensure we keep some connections alive
	}
	
	// Increase timeouts for external connections
	config.HealthCheckPeriod = 60 * time.Second
	config.ConnConfig.ConnectTimeout = 10 * time.Second
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 15 * time.Minute
	
	config.AfterConnect = afterConnectRegisterTypes

	// Add retry logic for initial connection
	var pool *pgxpool.Pool
	for retries := 0; retries < 5; retries++ {
		pool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/5): %v\n", retries+1, err)
		time.Sleep(2 * time.Second)
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after 5 attempts: %v", err)
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &DbClient{pool}, nil
}
