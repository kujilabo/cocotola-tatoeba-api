package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"

	"github.com/kujilabo/cocotola-tatoeba-api/pkg/handler/converter"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg/handler/entity"
	handlerhelper "github.com/kujilabo/cocotola-tatoeba-api/pkg/handler/helper"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg/usecase"
	lib "github.com/kujilabo/cocotola-tatoeba-api/pkg_lib/domain"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg_lib/ginhelper"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg_lib/log"
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
// @Summary     import links
// @Description import links
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
	handlerhelper.HandleFunction(c, func() error {
		param := entity.TatoebaSentenceFindParameter{}
		if err := c.ShouldBindJSON(&param); err != nil {
			c.Status(http.StatusBadRequest)
			return nil
		}
		parameter, err := converter.ToTatoebaSentenceSearchCondition(ctx, &param)
		if err != nil {
			return err
		}
		result, err := h.userUsecase.FindSentencePairs(ctx, parameter)
		if err != nil {
			return xerrors.Errorf("failed to FindSentences. err: %w", err)
		}
		response, err := converter.ToTatoebaSentenceFindResponse(ctx, result)
		if err != nil {
			return err
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
			return xerrors.Errorf("failed to FindSentences. err: %w", err)
		}
		response, err := converter.ToTatoebaSentenceResponse(ctx, result)
		if err != nil {
			return err
		}

		c.JSON(http.StatusOK, response)
		return nil
	}, h.errorHandle)
}

func (h *userHandler) errorHandle(c *gin.Context, err error) bool {
	ctx := c.Request.Context()
	logger := log.FromContext(ctx)
	logger.Errorf("userHandler. err: %v", err)
	return false
}
