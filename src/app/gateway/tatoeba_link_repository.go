package gateway

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/kujilabo/cocotola-tatoeba-api/src/app/service"
	liberrors "github.com/kujilabo/cocotola-tatoeba-api/src/lib/errors"
	libG "github.com/kujilabo/cocotola-tatoeba-api/src/lib/gateway"
)

type tatoebaLinkRepository struct {
	db           *gorm.DB
	sentenceRepo service.TatoebaSentenceRepository
}

type tatoebaLinkEntity struct {
	From int
	To   int
}

func (e *tatoebaLinkEntity) TableName() string {
	return "tatoeba_link"
}

func NewTatoebaLinkRepository(db *gorm.DB) (service.TatoebaLinkRepository, error) {
	sentenceRepo, err := NewTatoebaSentenceRepository(db)
	if err != nil {
		return nil, err
	}

	return &tatoebaLinkRepository{
		db:           db,
		sentenceRepo: sentenceRepo,
	}, nil
}

func (r *tatoebaLinkRepository) Add(ctx context.Context, param service.TatoebaLinkAddParameter) error {

	fromContained, err := r.sentenceRepo.ContainsSentenceBySentenceNumber(ctx, param.GetFrom())
	if err != nil {
		return err
	}

	toContained, err := r.sentenceRepo.ContainsSentenceBySentenceNumber(ctx, param.GetTo())
	if err != nil {
		return err
	}

	if !fromContained || !toContained {
		return service.ErrTatoebaSentenceNotFound
	}

	entity := tatoebaLinkEntity{
		From: param.GetFrom(),
		To:   param.GetTo(),
	}

	if result := r.db.Create(&entity); result.Error != nil {
		if err := libG.ConvertDuplicatedError(result.Error, service.ErrTatoebaLinkAlreadyExists); errors.Is(err, service.ErrTatoebaLinkAlreadyExists) {
			return liberrors.Errorf("failed to Add tatoebaLink. err: %w", err)
		}

		if err := libG.ConvertRelationError(result.Error, service.ErrTatoebaLinkSourceNotFound); errors.Is(err, service.ErrTatoebaLinkSourceNotFound) {
			fmt.Printf("relation %v, %v\n", fromContained, toContained)
			// nothing
			return nil
		}

		return liberrors.Errorf("failed to Add tatoebaLink. err: %w", result.Error)
	}

	return nil
}
