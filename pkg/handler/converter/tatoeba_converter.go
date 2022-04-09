package converter

import (
	"context"

	"github.com/kujilabo/cocotola-tatoeba-api/pkg/handler/entity"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg/service"
)

func ToTatoebaSentenceSearchCondition(ctx context.Context, param *entity.TatoebaSentenceFindParameter) (service.TatoebaSentenceSearchCondition, error) {
	return service.NewTatoebaSentenceSearchCondition(param.PageNo, param.PageSize, param.Keyword, param.Random)
}

func ToTatoebaSentenceFindResponse(ctx context.Context, result service.TatoebaSentencePairSearchResult) (*entity.TatoebaSentencePairFindResponse, error) {
	entities := make([]entity.TatoebaSentencePair, len(result.GetResults()))
	for i, m := range result.GetResults() {
		src := entity.TatoebaSentenceResponse{
			SentenceNumber: m.GetSrc().GetSentenceNumber(),
			Lang:           m.GetSrc().GetLang().String(),
			Text:           m.GetSrc().GetText(),
			Author:         m.GetSrc().GetAuthor(),
			UpdatedAt:      m.GetSrc().GetUpdatedAt(),
		}
		dst := entity.TatoebaSentenceResponse{
			SentenceNumber: m.GetDst().GetSentenceNumber(),
			Lang:           m.GetDst().GetLang().String(),
			Text:           m.GetDst().GetText(),
			Author:         m.GetDst().GetAuthor(),
			UpdatedAt:      m.GetDst().GetUpdatedAt(),
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
	return &entity.TatoebaSentenceResponse{
		SentenceNumber: result.GetSentenceNumber(),
		Lang:           result.GetLang().String(),
		Text:           result.GetText(),
		Author:         result.GetAuthor(),
		UpdatedAt:      result.GetUpdatedAt(),
	}, nil
}
