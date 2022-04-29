package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"

	handlerhelper "github.com/kujilabo/cocotola-tatoeba-api/src/app/handler/helper"
	"github.com/kujilabo/cocotola-tatoeba-api/src/app/service"
	"github.com/kujilabo/cocotola-tatoeba-api/src/app/usecase"
	"github.com/kujilabo/cocotola-tatoeba-api/src/lib/log"
)

type AdminHandler interface {
	ImportSentences(c *gin.Context)
	ImportLinks(c *gin.Context)
}

type adminHandler struct {
	adminUsecase                         usecase.AdminUsecase
	newTatoebaSentenceAddParameterReader func(reader io.Reader) service.TatoebaSentenceAddParameterIterator
	newTatoebaLinkAddParameterReader     func(reader io.Reader) service.TatoebaLinkAddParameterIterator
}

func NewAdminHandler(adminUsecase usecase.AdminUsecase, newTatoebaSentenceAddParameterReader func(reader io.Reader) service.TatoebaSentenceAddParameterIterator, newTatoebaLinkAddParameterReader func(reader io.Reader) service.TatoebaLinkAddParameterIterator) AdminHandler {
	return &adminHandler{
		adminUsecase:                         adminUsecase,
		newTatoebaSentenceAddParameterReader: newTatoebaSentenceAddParameterReader,
		newTatoebaLinkAddParameterReader:     newTatoebaLinkAddParameterReader,
	}
}

// ImportSentences godoc
// @Summary     import sentences
// @Description import sentences
// @Tags        tatoeba
// @Param       file formData file true "***_sentences_detailed.tsv"
// @Success     200
// @Failure     400
// @Failure     401
// @Failure     500
// @Router      /v1/admin/sentence/import [post]
// @Security    BasicAuth
func (h *adminHandler) ImportSentences(c *gin.Context) {
	ctx := c.Request.Context()
	logger := log.FromContext(ctx)
	handlerhelper.HandleFunction(c, func() error {
		file, err := c.FormFile("file")
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				logger.Warnf("err: %+v", err)
				c.Status(http.StatusBadRequest)
				return nil
			}
			if errors.Is(err, http.ErrNotMultipart) {
				logger.Warnf("err: %+v", err)
				c.Status(http.StatusBadRequest)
				return nil
			}
			return err
		}

		multipartFile, err := file.Open()
		if err != nil {
			return xerrors.Errorf("failed to file.Open. err: %w", err)
		}
		defer multipartFile.Close()

		iterator := h.newTatoebaSentenceAddParameterReader(multipartFile)

		if err := h.adminUsecase.ImportSentences(ctx, iterator); err != nil {
			return xerrors.Errorf("failed to ImportSentences. err: %w", err)
		}

		c.Status(http.StatusOK)
		return nil
	}, h.errorHandle)
}

// ImportLinks godoc
// @Summary     import links
// @Description import links
// @Tags        tatoeba
// @Param       file formData file true "links.csv"
// @Success     200
// @Failure     400
// @Failure     401
// @Failure     500
// @Router      /v1/admin/link/import [post]
// @Security    BasicAuth
func (h *adminHandler) ImportLinks(c *gin.Context) {
	ctx := c.Request.Context()
	logger := log.FromContext(ctx)
	handlerhelper.HandleFunction(c, func() error {
		file, err := c.FormFile("file")
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				logger.Warnf("err: %+v", err)
				c.Status(http.StatusBadRequest)
				return nil
			}
			if errors.Is(err, http.ErrNotMultipart) {
				logger.Warnf("err: %+v", err)
				c.Status(http.StatusBadRequest)
				return nil
			}
			return err
		}

		multipartFile, err := file.Open()
		if err != nil {
			return xerrors.Errorf("failed to file.Open. err: %w", err)
		}
		defer multipartFile.Close()

		iterator := h.newTatoebaLinkAddParameterReader(multipartFile)

		if err := h.adminUsecase.ImportLinks(ctx, iterator); err != nil {
			return xerrors.Errorf("failed to ImportLinks. err: %w", err)
		}

		c.Status(http.StatusOK)
		return nil
	}, h.errorHandle)
}

func (h *adminHandler) errorHandle(c *gin.Context, err error) bool {
	ctx := c.Request.Context()
	logger := log.FromContext(ctx)
	logger.Errorf("adminHandler. err: %v", err)
	return false
}
