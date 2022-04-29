package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kujilabo/cocotola-tatoeba-api/src/app/handler/entity"
	libD "github.com/kujilabo/cocotola-tatoeba-api/src/lib/domain"
)

func Test_TatoebaSentenceResponse_validation(t *testing.T) {
	tests := []struct {
		name       string
		entity     entity.TatoebaSentenceResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "lang2 is 'en'",
			entity: entity.TatoebaSentenceResponse{
				Lang2: "en",
			},
			wantErr: false,
		},
		{
			name: "lang2 is 'ja'",
			entity: entity.TatoebaSentenceResponse{
				Lang2: "ja",
			},
			wantErr: false,
		},
		{
			name: "lang2 is 'es'",
			entity: entity.TatoebaSentenceResponse{
				Lang2: "es",
			},
			wantErr:    true,
			wantErrMsg: "Key: 'TatoebaSentenceResponse.Lang2' Error:Field validation for 'Lang2' failed on the 'oneof' tag",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := libD.Validator.Struct(tt.entity)
			if tt.wantErr {
				assert.Equal(t, err.Error(), tt.wantErrMsg)
			}
		})
	}
}
