package domain

import (
	"errors"
	"strings"
	"time"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "PENDING"
)

type Payment struct {
	ID        string        `json:"id"`
	MerchantID string       `json:"merchant_id"`
	Amount    int64         `json:"amount"`
	Currency  string        `json:"currency"`
	Status    PaymentStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type CreatePaymentRequest struct {
	MerchantID string `json:"merchant_id"`
	Amount     int64  `json:"amount"`
	Currency   string `json:"currency"`
}

func (r CreatePaymentRequest) Validate() error {
	if strings.TrimSpace(r.MerchantID) == "" {
		return errors.New("merchant_id is required")
	}
	if r.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if len(strings.TrimSpace(r.Currency)) != 3 {
		return errors.New("currency must be a 3-letter ISO code")
	}
	return nil
}
