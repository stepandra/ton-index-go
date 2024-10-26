# TON Index Setup Guide

## Prerequisites

1. Install required software:
```bash
# Update package lists
sudo apt update

# Install PostgreSQL 15
sudo apt install -y postgresql-15

# Install Go 1.20+
wget https://go.dev/dl/go1.20.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

2. Configure PostgreSQL:
```bash
# Create database
sudo -u postgres createdb mainnet

# Create required extensions (connect to DB first)
sudo -u postgres psql mainnet -c "CREATE EXTENSION IF NOT EXISTS pg_trgm;"
```

3. Clone repository:
```bash 
git clone https://github.com/yourusername/ton-index-go.git
cd ton-index-go
```

## Local Ansible Setup

1. Create private config directory:
```bash
mkdir -p private
```

2. Create inventory file for localhost:
private/inventory.yaml:
```yaml
all:
  hosts:
    localhost:
      ansible_connection: local
```

3. Create config file:
private/config.yaml:
```yaml
---
application_name: ton-index-go
bind_addr: ":4100"
pg_dsn: postgresql://postgres@localhost:5432/mainnet
additional_args: ""
remote_hosts:
  - localhost
service_user: your_username  # Replace with your actual username
service_group: your_username # Replace with your actual username
go_binary: ton-index-go
go_binary_path: /usr/local/bin
systemd_service_path: /etc/systemd/system
```

## Deploy

Run ansible playbook:
```bash
ansible-playbook -i private/inventory.yaml systemd-deploy.yaml
```

## Add Bonding Curve Memepads Endpoint

To track specific contracts, we need to:

1. Create new endpoint `/api/v3/memepads/bondingCurve`
2. Add contract tracking logic
3. Store contract states in database

Required SQL migrations:

```sql
-- Create table for bonding curve states
CREATE TABLE IF NOT EXISTS bonding_curve_states (
    address text PRIMARY KEY,
    token_balance numeric(20,0),
    ton_balance numeric(20,0),
    last_price numeric(20,0),
    last_updated timestamp with time zone,
    last_tx_lt bigint
);

-- Create index on last_updated
CREATE INDEX IF NOT EXISTS idx_bonding_curve_last_updated 
ON bonding_curve_states(last_updated);
```

The endpoint will return current state of bonding curve contracts including:
- Token balance
- TON balance  
- Last price
- Last update timestamp

See code changes below for implementation details.
