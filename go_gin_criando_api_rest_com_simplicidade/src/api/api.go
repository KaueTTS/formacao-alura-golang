package api

import (
	healthRoute "go_gin_criando_api_rest_com_simplicidade/src/api/routes/health"
	studentRoute "go_gin_criando_api_rest_com_simplicidade/src/api/routes/student"
	postgresdb "go_gin_criando_api_rest_com_simplicidade/src/config/db"
	"go_gin_criando_api_rest_com_simplicidade/src/config/env"
	"net/http"

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
