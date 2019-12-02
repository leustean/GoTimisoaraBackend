package middlewares

import (
	"github.com/gin-gonic/gin"
	"goTimisoaraBackend/auth"
)

import (
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) func(c *gin.Context, w http.ResponseWriter, r *http.Request) {
	return func(c *gin.Context, w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		next(w, r)
	}
}
