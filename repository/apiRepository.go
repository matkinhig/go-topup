package repository

import "github.com/matkinhig/go-topup/models"

type ApiRepository interface {
	GetAll() ([]models.DepositRepository, error)
	// GetByAwardCode() ([]models.DepositRepository, error)
	// CreateDeposit() ([]models.Response, error)
	// UpdateDeposit() ([]models.Response, error)
}
