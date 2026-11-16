package service

import (
	"context"
	"errors"

	"payflow/internal/payment/domain"
	"payflow/internal/payment/repository"
)

type Service struct {
	repo repository.PaymentRepository
}

func NewService(repo repository.PaymentRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePayment(ctx context.Context, req domain.CreatePaymentRequest) (*domain.Payment, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	payment := &domain.Payment{
		MerchantID: req.MerchantID,
		Amount:     req.Amount,
		Currency:   req.Currency,
	}

	if err := s.repo.Create(ctx, payment); err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *Service) GetPayment(ctx context.Context, id string) (*domain.Payment, error) {
	if id == "" {
		return nil, errors.New("payment id is required")
	}
	return s.repo.GetByID(ctx, id)
}
