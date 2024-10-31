package main

import "log"

// generateEmbedding creates a vector embedding for a given text.
func generateEmbedding(text string) ([]float64, error) {
    var embedding []float64
    query := `SELECT pgai_generate_embedding($1);`
    if err := Db.Raw(query, text).Scan(&embedding).Error; err != nil {
        log.Printf("Error generating embedding: %v", err)
        return nil, err
    }
    return embedding, nil
}
