package routes

import (
	"awesomeProject/api"
	"awesomeProject/db"
	"github.com/gin-gonic/gin"
)

func SetupAPIHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiHandler := api.NewHandler(c, queries)
		c.Set("apiHandler", apiHandler)
		c.Next()
	}
}
