package router

import (
	"github.com/gin-gonic/gin"
	"api/db"
	"api/handler"
)

func SetupRouter(conn *db.ConnWrapper) *gin.Engine {
	r := gin.Default()

	r.Get("/keywards", handler.GetKeywardsHandler(func() *db.ConnWrapper {
		return conn
	}))

	r.Post("/keywards", handler.InsertKeywardsHandler(func() *db.ConnWrapper {
		return conn
	}))

	return r
}
