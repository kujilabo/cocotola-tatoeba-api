package usecase

import (
	"context"

	"gorm.io/gorm"

	"github.com/kujilabo/cocotola-tatoeba-api/src/app/service"
	liberrors "github.com/kujilabo/cocotola-tatoeba-api/src/lib/errors"
)

type UserUsecase interface {
	FindSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchCondition) (service.TatoebaSentencePairSearchResult, error)

	FindSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (service.TatoebaSentence, error)
}

type userUsecase struct {
	db     *gorm.DB
	rfFunc service.RepositoryFactoryFunc
}

func NewUserUsecase(db *gorm.DB, rfFunc service.RepositoryFactoryFunc) UserUsecase {
	return &userUsecase{
		db:     db,
		rfFunc: rfFunc,
	}
}

func (u *userUsecase) FindSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchCondition) (service.TatoebaSentencePairSearchResult, error) {
	var result service.TatoebaSentencePairSearchResult
	if err := u.db.Transaction(func(tx *gorm.DB) error {
		rf, err := u.rfFunc(ctx, tx)
		if err != nil {
			return liberrors.Errorf("create RepositoryFactory. err: %w", err)
		}

		repo, err := rf.NewTatoebaSentenceRepository(ctx)
		if err != nil {
			return liberrors.Errorf("new TatoebaSentenceRepository. err: %w", err)
		}

		tmpResult, err := repo.FindTatoebaSentencePairs(ctx, param)
		if err != nil {
			return liberrors.Errorf("execute FindTatoebaSentencePairs. err: %w", err)
		}
		result = tmpResult
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userUsecase) FindSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (service.TatoebaSentence, error) {
	var result service.TatoebaSentence
	if err := u.db.Transaction(func(tx *gorm.DB) error {
		rf, err := u.rfFunc(ctx, tx)
		if err != nil {
			return liberrors.Errorf("create RepositoryFactory. err: %w", err)
		}

		repo, err := rf.NewTatoebaSentenceRepository(ctx)
		if err != nil {
			return liberrors.Errorf("new TatoebaSentenceRepository. err: %w", err)
		}

		tmpResult, err := repo.FindTatoebaSentenceBySentenceNumber(ctx, sentenceNumber)
		if err != nil {
			return liberrors.Errorf("execute FindTatoebaSentenceBySentenceNumber. err: %w", err)
		}
		result = tmpResult
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}
