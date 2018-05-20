package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"os"
)
var SSL_HOST = os.Getenv("HOST")+":"+os.Getenv("HTTPS_PORT")
var SecureMiddlewareConfig = secure.New(secure.Options{
	SSLRedirect: true,
	SSLHost: SSL_HOST,
})

var RedirectTo = func() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := SecureMiddlewareConfig.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			c.Abort()
			return
		}
	}
}()