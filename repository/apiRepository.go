package repository

import "github.com/matkinhig/go-topup/models"

type ApiRepository interface {
	FindByCustID(*models.RequestGet) (*models.ResponseGet, error)
	// GetByAwardCode() ([]models.DepositRepository, error)
	// CreateDeposit() ([]models.Response, error)
	// UpdateDeposit() ([]models.Response, error)
}
