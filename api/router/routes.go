package router

import (
	"api/handler"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupRouter wires the keyword routes.
func SetupRouter(pool *pgxpool.Pool) *gin.Engine {
	r := gin.Default()
	r.GET("/keywords", handler.GetKeywordsHandler(pool))
	r.POST("/keywords", handler.InsertKeywordHandler(pool))
	return r
}
