package usecase

import (
	"context"
	"errors"
	"io"

	"gorm.io/gorm"

	"github.com/kujilabo/cocotola-tatoeba-api/src/app/service"
	liberrors "github.com/kujilabo/cocotola-tatoeba-api/src/lib/errors"
	"github.com/kujilabo/cocotola-tatoeba-api/src/lib/log"
)

const (
	commitSize = 1000
	logSize    = 100000
)

type AdminUsecase interface {
	ImportSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator) error

	ImportLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator) error
}

type adminUsecase struct {
	db     *gorm.DB
	rfFunc service.RepositoryFactoryFunc
}

func NewAdminUsecase(db *gorm.DB, rfFunc service.RepositoryFactoryFunc) AdminUsecase {
	return &adminUsecase{
		db:     db,
		rfFunc: rfFunc,
	}
}

func (u *adminUsecase) ImportSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator) error {
	logger := log.FromContext(ctx)

	var readCount = 0
	var importCount = 0
	var skipCount = 0
	var loop = true
	for loop {
		if err := u.db.Transaction(func(tx *gorm.DB) error {
			rf, err := u.rfFunc(ctx, tx)
			if err != nil {
				return liberrors.Errorf("create RepositoryFactory. err: %w", err)
			}

			repo, err := rf.NewTatoebaSentenceRepository(ctx)
			if err != nil {
				return liberrors.Errorf("new TatoebaSentenceRepository. err: %w", err)
			}

			i := 0
			for {
				param, err := iterator.Next(ctx)
				if errors.Is(err, io.EOF) {
					loop = false
					break
				}
				readCount++
				if err != nil {
					return liberrors.Errorf("read next line. read count: %d, err: %w", readCount, err)
				}

				if param == nil {
					skipCount++
					continue
				}

				if err := repo.Add(ctx, param); err != nil {
					logger.Warnf("failed to Add. read count: %d, err: %v", readCount, err)
					skipCount++
					continue
				}

				i++
				importCount++
				if i >= commitSize {
					if importCount%logSize == 0 {
						logger.Infof("imported count: %d", importCount)
					}
					break
				}
			}

			return nil
		}); err != nil {
			return liberrors.Errorf("import sentence. err: %w", err)
		}
	}

	logger.Infof("imported count: %d", importCount)
	logger.Infof("skipped count: %d", skipCount)
	logger.Infof("read count: %d", readCount)

	return nil
}

func (u *adminUsecase) ImportLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator) error {
	logger := log.FromContext(ctx)

	var readCount = 0
	var importCount = 0
	var skipCount = 0
	var loop = true
	for loop {
		if err := u.db.Transaction(func(tx *gorm.DB) error {
			rf, err := u.rfFunc(ctx, tx)
			if err != nil {
				return liberrors.Errorf("create RepositoryFactory. err: %w", err)
			}

			repo, err := rf.NewTatoebaLinkRepository(ctx)
			if err != nil {
				return liberrors.Errorf("new TatoebaLinkRepository. err: %w", err)
			}

			i := 0
			for {
				param, err := iterator.Next(ctx)
				if errors.Is(err, io.EOF) {
					loop = false
					break
				}
				readCount++
				if err != nil {
					return liberrors.Errorf("read next line. read count: %d, err: %w", readCount, err)
				}
				if param == nil {
					skipCount++
					continue
				}

				if err := repo.Add(ctx, param); err != nil {
					if !errors.Is(err, service.ErrTatoebaSentenceNotFound) {
						logger.Warnf("failed to Add. read count: %d, err: %v", readCount, err)
					}
					skipCount++
					continue
				}
				i++
				importCount++
				if i >= commitSize {
					if importCount%logSize == 0 {
						logger.Infof("imported count: %d", importCount)
					}
					break
				}
			}

			return nil
		}); err != nil {
			return liberrors.Errorf("import sentence. err: %w", err)
		}
	}

	logger.Infof("imported count: %d", importCount)
	logger.Infof("skipped count: %d", skipCount)
	logger.Infof("read count: %d", readCount)

	return nil
}
