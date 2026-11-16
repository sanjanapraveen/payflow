# Architecture

## Current scope
- Go HTTP API
- PostgreSQL persistence
- payment create and retrieval
- validation
- health endpoints

## Components
- `cmd/api` starts the HTTP server
- `internal/config` loads environment settings
- `internal/payment/domain` holds payment and validation types
- `internal/payment/repository` persists payments
- `internal/payment/service` contains the business logic
- `internal/transport/http` maps HTTP requests to the service

## Planned next steps
- idempotency
- ledger invariants
- provider simulator
- outbox + event workers
- webhooks + reconciliation
