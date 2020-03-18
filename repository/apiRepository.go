package repository

import "github.com/matkinhig/go-topup/models"

type ApiRepository interface {
	FindByCustID(*models.RequestGet) (*models.ResponseQuery, error)
	CreateDeposit(*models.RequestPost) (*models.ResponseCreate, error)
	UpdateDeposit(*models.RequestUpdate) (*models.ResponseUpdate, error)
}
