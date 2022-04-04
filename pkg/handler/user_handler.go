package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"

	"github.com/kujilabo/cocotola-tatoeba-api/pkg/handler/converter"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg/handler/entity"
	handlerhelper "github.com/kujilabo/cocotola-tatoeba-api/pkg/handler/helper"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg/usecase"
	"github.com/kujilabo/cocotola-tatoeba-api/pkg_lib/log"
)

type UserHandler interface {
	FindSentences(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) FindSentences(c *gin.Context) {
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
		result, err := h.userUsecase.FindSentences(ctx, parameter)
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
