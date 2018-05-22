package middlewares

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

func RevisionMiddleware() gin.HandlerFunc {
	data, err := ioutil.ReadFile("REVISION")
	if err != nil {
		return func(c *gin.Context) {
			c.Next()
		}
	}
	revision := strings.TrimSpace(string(data))

	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Revision", revision)
		c.Next()
	}
}