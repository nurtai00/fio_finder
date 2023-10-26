package repository

import (
	"context"
	"fio_finder/internal/models"
)

//go:generate mockgen -source=person.go -destination=mocks/person.go
type PersonRepository interface {
	Create(ctx context.Context, person *models.Person) error
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, id uint64, fieldsToUpdate models.PersonFieldsToUpdate) error
	Get(ctx context.Context, id uint64) (*models.Person, error)
	GetList(ctx context.Context) ([]models.Person, error)
}
