package gateway

import (
	"context"

	"gorm.io/gorm"

	"github.com/kujilabo/cocotola-tatoeba-api/src/app/service"
)

type repositoryFactory struct {
	db         *gorm.DB
	driverName string
}

func NewRepositoryFactory(ctx context.Context, db *gorm.DB, driverName string) (service.RepositoryFactory, error) {
	return &repositoryFactory{
		db:         db,
		driverName: driverName,
	}, nil
}

func (f *repositoryFactory) NewTatoebaSentenceRepository(ctx context.Context) (service.TatoebaSentenceRepository, error) {
	return NewTatoebaSentenceRepository(f.db)
}

func (f *repositoryFactory) NewTatoebaLinkRepository(ctx context.Context) (service.TatoebaLinkRepository, error) {
	return NewTatoebaLinkRepository(f.db)
}
