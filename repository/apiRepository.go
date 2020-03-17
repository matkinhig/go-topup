package repository

import "github.com/matkinhig/go-topup/models"

type ApiRepository interface {
	FindByCustID(*models.RequestGet) (*models.ResponseQuery, error)
	// GetByAwardCode() ([]models.DepositRepository, error)
	CreateDeposit(*models.RequestPost) (*models.ResponseCreate, error)
	// UpdateDeposit() ([]models.Response, error)
}
