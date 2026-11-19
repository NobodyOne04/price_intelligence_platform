package handler

import (
	"net/http"
	"strings"

	"api/db"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// GetKeywordsHandler lists the enabled keywords.
func GetKeywordsHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		keywords, err := db.GetKeywords(c.Request.Context(), pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get keywords"})
			return
		}
		c.JSON(http.StatusOK, keywords)
	}
}

type newKeyword struct {
	Value string `json:"value" binding:"required"`
}

// InsertKeywordHandler adds a keyword from {"value": "..."}.
func InsertKeywordHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var k newKeyword
		if err := c.ShouldBindJSON(&k); err != nil || strings.TrimSpace(k.Value) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input: want {\"value\": \"...\"}"})
			return
		}
		if err := db.InsertKeyword(c.Request.Context(), pool, strings.TrimSpace(k.Value)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert keyword"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "keyword added"})
	}
}
