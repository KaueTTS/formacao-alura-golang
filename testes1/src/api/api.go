package api

import (
	"net/http"
	healthRoute "testes1/src/api/routes/health"
	studentRoute "testes1/src/api/routes/student"
	postgresdb "testes1/src/config/db"
	"testes1/src/config/env"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	cfg := env.Load()
	postgresdb.Connect(cfg)
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route_not_found"})
	})
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method_not_allowed"})
	})

	InjectRoutes(r)

	return r
}

func InjectRoutes(app *gin.Engine) {
	healthRoute.HealthRoute(app)
	studentRoute.StudentRoutes(app)
}
