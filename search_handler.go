package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchHandler processes search queries and returns relevant papers.
func SearchHandler(c *gin.Context) {
	var req struct {
		Query string `json:"query"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	embedding, err := generateEmbedding(req.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate embedding"})
		return
	}

	var results []Paper
	query := `
		SELECT *, 1 - (embedding <=> $1) AS similarity
		FROM papers
		ORDER BY embedding <=> $1
		LIMIT 10;
	`
	err = Db.Select(&results, query, embedding)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return
	}

	c.JSON(http.StatusOK, results)
}
