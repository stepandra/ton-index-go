# Utils

This directory contains utility programs and scripts for TON Index GO.

## Database Setup

Before running the API server, you'll need to:

1. Create a PostgreSQL database
2. Initialize the schema by running the SQL migration files:
```bash
psql -U <user> -d <dbname> -f migrations/01_init.sql
psql -U <user> -d <dbname> -f migrations/02_nft_tables.sql
psql -U <user> -d <dbname> -f migrations/03_jetton_tables.sql
```

## Utilities

### fix-broken-traces

A utility for fixing broken traces in the database. Only needed if you encounter trace consistency issues.

Usage:
```bash
go run utils/fix-broken-traces/main.go \
  --pg "postgresql://user:pass@localhost:5432/dbname" \
  --batch 100 \
  --processes 32
```

Options:
- `--pg`: PostgreSQL connection string
- `--batch`: Number of traces to process in each batch (default: 100) 
- `--processes`: Number of parallel workers (default: 32)
- `--total`: Override total trace count
