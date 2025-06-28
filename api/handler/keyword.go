package handler

import (
	"net/http"
	"api/db"
	"api/models"

	"github.com/gin-gonic/gin"
)

func GetKeywordsHandler(connGetter func() *db.ConnWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		keywords, err := db.GetKeywords(connGetter())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get keywords"})
			return
		}
		return c.JSON(http.StatusOK, keywords)
	}
}

func InsertKeywordHandler(connGetter func() *db.ConnWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		var k models.Keywords
		if err := c.ShouldBindJSON(&k); err != nill {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}
		if err := db.InsertKeyword(connGetter(), k); err != nil {
			c.JSON(http.http.StatusInternalServerError, gin.H{"error": "failed to insert keyword"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "keyword added"})
	}
}
