package handlerhelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFunction(c *gin.Context, fn func() error, errorHandle func(c *gin.Context, err error) bool) {
	if err := fn(); err != nil {
		if handled := errorHandle(c, err); !handled {
			c.Status(http.StatusInternalServerError)
		}
	}
}
