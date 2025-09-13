package route

import (
	"github.com/gin-gonic/gin"
)

type Health struct {
	Status string `json:"status"`
}

func HealthRoute(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, Health{Status: "ok"})
	})
}
