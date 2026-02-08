package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"time"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) GetTodaySummary() (*models.SalesReport, error) {
	return s.repo.GetTodaySummary()
}

func (s *TransactionService) GetSummaryByDateRange(startDate, endDate time.Time) (*models.SalesReport, error) {
	return s.repo.GetSummaryByDateRange(startDate, endDate)
}
