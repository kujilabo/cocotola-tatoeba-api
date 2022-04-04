package converter

import (
	"context"

	"github.com/kujilabo/cocotola-tatoeba-api/pkg/handler/entity"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg/service"
)

func ToTatoebaSentenceSearchCondition(ctx context.Context, param *entity.TatoebaSentenceFindParameter) (service.TatoebaSentenceSearchCondition, error) {
	return service.NewTatoebaSentenceSearchCondition(param.PageNo, param.PageSize, param.Keyword, param.Random)
}

func ToTatoebaSentenceResponse(ctx context.Context, result *service.TatoebaSentenceSearchResult) (*entity.TatoebaSentenceFindResponse, error) {
	entities := make([]entity.TatoebaSentencePair, len(result.Results))
	for i, m := range result.Results {
		src := entity.TatoebaSentence{
			SentenceNumber: m.GetSrc().GetSentenceNumber(),
			Lang:           m.GetSrc().GetLang().String(),
			Text:           m.GetSrc().GetText(),
			Author:         m.GetSrc().GetAuthor(),
			UpdatedAt:      m.GetSrc().GetUpdatedAt(),
		}
		dst := entity.TatoebaSentence{
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

	return &entity.TatoebaSentenceFindResponse{
		TotalCount: result.TotalCount,
		Results:    entities,
	}, nil
}
