package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"

	"payflow/internal/payment/domain"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *domain.Payment) error
	GetByID(ctx context.Context, id string) (*domain.Payment, error)
}

type PostgresPaymentRepository struct {
	db *sql.DB
}

func NewPostgresPaymentRepository(databaseURL string) (*PostgresPaymentRepository, error) {
	if databaseURL == "" {
		return nil, errors.New("database url is required")
	}

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return &PostgresPaymentRepository{db: db}, nil
}

func (r *PostgresPaymentRepository) Close() error {
	return r.db.Close()
}

func (r *PostgresPaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	payment.ID = uuid.NewString()
	payment.Status = domain.PaymentStatusPending
	now := time.Now().UTC()
	payment.CreatedAt = now
	payment.UpdatedAt = now

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO payments (id, merchant_id, amount, currency, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, payment.ID, payment.MerchantID, payment.Amount, payment.Currency, payment.Status, payment.CreatedAt, payment.UpdatedAt)
	return err
}

func (r *PostgresPaymentRepository) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, merchant_id, amount, currency, status, created_at, updated_at
		FROM payments
		WHERE id = $1
	`, id)

	p := &domain.Payment{}
	err := row.Scan(&p.ID, &p.MerchantID, &p.Amount, &p.Currency, &p.Status, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}

	return p, nil
}
