package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, dbConnection *sql.DB) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HEALTHCHECK -> OK",
		})
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "API/V1 -> OK",
			})
		})

		UsersRoutes(apiV1, dbConnection)
	}
}
