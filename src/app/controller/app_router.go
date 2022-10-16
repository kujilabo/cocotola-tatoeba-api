package controller

import (
	"io"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginlog "github.com/onrik/logrus/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/kujilabo/cocotola-tatoeba-api/src/app/config"
	"github.com/kujilabo/cocotola-tatoeba-api/src/app/gateway"
	"github.com/kujilabo/cocotola-tatoeba-api/src/app/service"
	"github.com/kujilabo/cocotola-tatoeba-api/src/app/usecase"
	"github.com/kujilabo/cocotola-tatoeba-api/src/lib/controller/middleware"
)

func NewRouter(adminUsecase usecase.AdminUsecase, userUsecase usecase.UserUsecase, corsConfig cors.Config, appConfig *config.AppConfig, authConfig *config.AuthConfig, debugConfig *config.DebugConfig) *gin.Engine {
	if !debugConfig.GinMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(cors.New(corsConfig))
	router.Use(gin.Recovery())

	if debugConfig.GinMode {
		router.Use(ginlog.Middleware(ginlog.DefaultConfig))
	}

	if debugConfig.Wait {
		router.Use(middleware.NewWaitMiddleware())
	}

	authMiddleware := gin.BasicAuth(gin.Accounts{
		authConfig.Username: authConfig.Password,
	})

	v1 := router.Group("v1")
	{
		v1.Use(otelgin.Middleware(appConfig.Name))
		v1.Use(middleware.NewTraceLogMiddleware(appConfig.Name))
		v1.Use(authMiddleware)
		{
			newSentenceReader := func(reader io.Reader) service.TatoebaSentenceAddParameterIterator {
				return gateway.NewTatoebaSentenceAddParameterReader(reader)
			}
			newLinkReader := func(reader io.Reader) service.TatoebaLinkAddParameterIterator {
				return gateway.NewTatoebaLinkAddParameterReader(reader)
			}

			admin := v1.Group("admin")
			adminHandler := NewAdminHandler(adminUsecase, newSentenceReader, newLinkReader)
			admin.POST("sentence/import", adminHandler.ImportSentences)
			admin.POST("link/import", adminHandler.ImportLinks)
		}
		{
			user := v1.Group("user")
			userHandler := NewUserHandler(userUsecase)
			user.POST("sentence_pair/find", userHandler.FindSentencePairs)
			user.GET("sentence/:sentenceNumber", userHandler.FindSentenceBySentenceNumber)
		}
	}

	return router
}
