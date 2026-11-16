CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL CHECK (amount > 0),
    currency VARCHAR(3) NOT NULL CHECK (currency ~ '^[A-Z]{3}$'),
    status VARCHAR(32) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_payments_merchant_id ON payments (merchant_id);
CREATE INDEX IF NOT EXISTS idx_payments_status ON payments (status);
