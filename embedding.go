package main

// generateEmbedding creates a vector embedding for a given text.
func generateEmbedding(text string) ([]float64, error) {
	var embedding []float64
	query := `SELECT pgai_generate_embedding($1);`
	err := Db.Get(&embedding, query, text)
	return embedding, err
}
