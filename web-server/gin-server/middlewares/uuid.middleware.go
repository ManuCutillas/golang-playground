package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"fmt"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidData, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("UUID Request Id Middleware Error: %s", err)
			return
		}
		c.Writer.Header().Set("X-Request-Id", uuidData.String())
		c.Next()
	}
}
