package converter

import (
	"context"

	"github.com/kujilabo/cocotola-tatoeba-api/pkg/handler/entity"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg/service"
	libD "github.com/kujilabo/cocotola-tatoeba-api/pkg_lib/domain"
)

func ToTatoebaSentenceSearchCondition(ctx context.Context, param *entity.TatoebaSentenceFindParameter) (service.TatoebaSentenceSearchCondition, error) {
	return service.NewTatoebaSentenceSearchCondition(param.PageNo, param.PageSize, param.Keyword, param.Random)
}

func ToTatoebaSentenceFindResponse(ctx context.Context, result service.TatoebaSentencePairSearchResult) (*entity.TatoebaSentencePairFindResponse, error) {
	entities := make([]entity.TatoebaSentencePair, len(result.GetResults()))
	for i, m := range result.GetResults() {
		src := entity.TatoebaSentenceResponse{
			SentenceNumber: m.GetSrc().GetSentenceNumber(),
			Lang2:          m.GetSrc().GetLang3().ToLang2().String(),
			Text:           m.GetSrc().GetText(),
			Author:         m.GetSrc().GetAuthor(),
			UpdatedAt:      m.GetSrc().GetUpdatedAt(),
		}
		if err := libD.Validator.Struct(src); err != nil {
			return nil, err
		}

		dst := entity.TatoebaSentenceResponse{
			SentenceNumber: m.GetDst().GetSentenceNumber(),
			Lang2:          m.GetDst().GetLang3().ToLang2().String(),
			Text:           m.GetDst().GetText(),
			Author:         m.GetDst().GetAuthor(),
			UpdatedAt:      m.GetDst().GetUpdatedAt(),
		}
		if err := libD.Validator.Struct(dst); err != nil {
			return nil, err
		}

		entities[i] = entity.TatoebaSentencePair{
			Src: src,
			Dst: dst,
		}
	}

	return &entity.TatoebaSentencePairFindResponse{
		TotalCount: result.GetTotalCount(),
		Results:    entities,
	}, nil
}

func ToTatoebaSentenceResponse(ctx context.Context, result service.TatoebaSentence) (*entity.TatoebaSentenceResponse, error) {
	e := &entity.TatoebaSentenceResponse{
		SentenceNumber: result.GetSentenceNumber(),
		Lang2:          result.GetLang3().ToLang2().String(),
		Text:           result.GetText(),
		Author:         result.GetAuthor(),
		UpdatedAt:      result.GetUpdatedAt(),
	}
	return e, libD.Validator.Struct(e)
}
