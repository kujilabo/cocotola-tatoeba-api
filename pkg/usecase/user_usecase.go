package usecase

import (
	"context"

	"gorm.io/gorm"

	"github.com/kujilabo/cocotola-tatoeba-api/pkg/service"
)

type UserUsecase interface {
	FindSentences(ctx context.Context, param service.TatoebaSentenceSearchCondition) (*service.TatoebaSentenceSearchResult, error)
}

type userUsecase struct {
	db *gorm.DB
	rf service.RepositoryFactoryFunc
}

func NewUserUsecase(db *gorm.DB, rf service.RepositoryFactoryFunc) UserUsecase {
	return &userUsecase{
		db: db,
		rf: rf,
	}
}

func (u *userUsecase) FindSentences(ctx context.Context, param service.TatoebaSentenceSearchCondition) (*service.TatoebaSentenceSearchResult, error) {
	var result *service.TatoebaSentenceSearchResult
	if err := u.db.Transaction(func(tx *gorm.DB) error {
		rf, err := u.rf(ctx, tx)
		if err != nil {
			return err
		}

		repo, err := rf.NewTatoebaSentenceRepository(ctx)
		if err != nil {
			return err
		}

		tmpResult, err := repo.FindTatoebaSentences(ctx, param)
		if err != nil {
			return err
		}
		result = tmpResult
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}
