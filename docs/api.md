# API

## GET /healthz
Returns basic service health.

Example:
```json
{"status":"ok"}
```

## GET /readyz
Returns readiness.

Example:
```json
{"status":"ready"}
```

## POST /v1/payments
Creates a payment.

Request:
```json
{
  "merchant_id": "merchant_123",
  "amount": 2500,
  "currency": "USD"
}
```

Response:
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "merchant_id": "merchant_123",
  "amount": 2500,
  "currency": "USD",
  "status": "PENDING",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

## GET /v1/payments/{id}
Returns a single payment by ID.
