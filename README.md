# TON Index GO

Fast and convenient TON Index API.

## Prerequisites

- Go 1.19 or later
- PostgreSQL 14 or later
- Git

## Local Development Setup

1. Clone the repository:
```bash
git clone https://github.com/kdimentionaltree/ton-index-go.git
cd ton-index-go
```

2. Install Go dependencies:
```bash
go mod download
```

3. Configure PostgreSQL:
- Create a new database
- Note your connection string (e.g. `postgresql://user:pass@localhost:5432/dbname`)

4. Run the service:
```bash
go run main.go \
  --pg "postgresql://user:pass@localhost:5432/dbname" \
  --bind ":8000" \
  --debug true
```

Available flags:
- `--pg`: PostgreSQL connection string (required)
- `--bind`: Server bind address (default ":8000")
- `--debug`: Enable debug mode
- `--maxconns`: PostgreSQL max connections (default 100)
- `--minconns`: PostgreSQL min connections (default 0)
- `--query-timeout`: Query timeout in milliseconds (default 3000)
- `--testnet`: Use testnet address book
- `--v2`: TON HTTP API endpoint for proxied methods
- `--v2-apikey`: API key for TON HTTP API endpoint

## Deploy with ansible

1. Copy `config.yaml` and `inventory.yaml` to `private/` folder.
2. Adjust `config.yaml` file.
3. Run `ansible-playbook -e private/config.yaml systemd-deploy.yaml`.

## API Documentation

Once running, visit http://localhost:8000/api/v3/ for Swagger UI documentation.
