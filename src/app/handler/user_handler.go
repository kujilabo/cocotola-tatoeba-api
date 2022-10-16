package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kujilabo/cocotola-tatoeba-api/src/app/handler/converter"
	"github.com/kujilabo/cocotola-tatoeba-api/src/app/handler/entity"
	handlerhelper "github.com/kujilabo/cocotola-tatoeba-api/src/app/handler/helper"
	"github.com/kujilabo/cocotola-tatoeba-api/src/app/usecase"
	lib "github.com/kujilabo/cocotola-tatoeba-api/src/lib/domain"
	liberrors "github.com/kujilabo/cocotola-tatoeba-api/src/lib/errors"
	"github.com/kujilabo/cocotola-tatoeba-api/src/lib/ginhelper"
	"github.com/kujilabo/cocotola-tatoeba-api/src/lib/log"
)

type UserHandler interface {
	FindSentencePairs(c *gin.Context)

	FindSentenceBySentenceNumber(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

// FindSentencePairs godoc
// @Summary     find pair of sentences
// @Description find pair of sentences
// @Tags        tatoeba
// @Accept      json
// @Produce     json
// @Param       param body entity.TatoebaSentenceFindParameter true "parameter to find sentences"
// @Success     200 {object} entity.TatoebaSentencePairFindResponse
// @Failure     400
// @Failure     401
// @Router      /v1/user/sentence_pair/find [post]
// @Security    BasicAuth
func (h *userHandler) FindSentencePairs(c *gin.Context) {
	ctx := c.Request.Context()
	logger := log.FromContext(ctx)

	handlerhelper.HandleFunction(c, func() error {
		param := entity.TatoebaSentenceFindParameter{}
		if err := c.ShouldBindJSON(&param); err != nil {
			c.Status(http.StatusBadRequest)
			return nil
		}
		logger.Debugf("FindSentencePairs. param: %+v", param)
		parameter, err := converter.ToTatoebaSentenceSearchCondition(ctx, &param)
		if err != nil {
			return liberrors.Errorf("convert parameter to TatoebaSentenceSearchCondition. err: %w", err)
		}
		result, err := h.userUsecase.FindSentencePairs(ctx, parameter)
		if err != nil {
			return liberrors.Errorf("execute FindSentencePairs. err: %w", err)
		}
		response, err := converter.ToTatoebaSentenceFindResponse(ctx, result)
		if err != nil {
			return liberrors.Errorf("convert result to TatoebaSentenceFindResponse. err: %w", err)
		}

		c.JSON(http.StatusOK, response)
		return nil
	}, h.errorHandle)
}

// FindSentenceBySentenceNumber godoc
// @Summary     import links
// @Description import links
// @Tags        tatoeba
// @Accept      json
// @Produce     json
// @Param       sentenceNumber path int true "Sentence number"
// @Success     200 {object} entity.TatoebaSentenceResponse
// @Failure     400
// @Failure     401
// @Router      /v1/user/sentence/{sentenceNumber} [get]
// @Security    BasicAuth
func (h *userHandler) FindSentenceBySentenceNumber(c *gin.Context) {
	ctx := c.Request.Context()
	handlerhelper.HandleFunction(c, func() error {
		sentenceNumber, err := ginhelper.GetIntFromPath(c, "sentenceNumber")
		if err != nil {
			return lib.ErrInvalidArgument
		}

		result, err := h.userUsecase.FindSentenceBySentenceNumber(ctx, sentenceNumber)
		if err != nil {
			return liberrors.Errorf("execute FindSentenceBySentenceNumber. err: %w", err)
		}
		response, err := converter.ToTatoebaSentenceResponse(ctx, result)
		if err != nil {
			return liberrors.Errorf("convert result to TatoebaSentenceResponse. err: %w", err)
		}

		c.JSON(http.StatusOK, response)
		return nil
	}, h.errorHandle)
}

func (h *userHandler) errorHandle(c *gin.Context, err error) bool {
	ctx := c.Request.Context()
	logger := log.FromContext(ctx)
	logger.Errorf("userHandler. err: %+v", err)
	return false
}
