# PayFlow

A minimal Go payment-processing backend for Milestone 0 and Milestone 1.

## Prerequisites
- Go 1.22+
- Docker + Docker Compose

## Local setup
1. Start PostgreSQL:
   ```bash
   docker compose up -d postgres
   ```
2. Apply the migration:
   ```bash
   psql "postgresql://payflow:payflow@localhost:5432/payflow" -f migrations/001_init.up.sql
   ```
3. Run the API:
   ```bash
   go run ./cmd/api
   ```

## Health endpoints
- GET /healthz
- GET /readyz

## Payment endpoints
- POST /v1/payments
- GET /v1/payments/{id}

Example create request:
```json
{
  "merchant_id": "merchant_123",
  "amount": 2500,
  "currency": "USD"
}
```

## Current scope
This project intentionally stops at:
- config + health checks
- Postgres persistence
- payment create + lookup
- validation
- local Docker-backed setup

This is Milestone 0 and Milestone 1 only. Financial correctness and concurrency-sensitive logic are intentionally deferred.