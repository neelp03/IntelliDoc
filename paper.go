package main

import (
    "time"
		"log"
)

type Paper struct {
    ID              int       `gorm:"primaryKey"`
    Title           string
    Authors         string
    PublicationDate time.Time
    Section         string
    Content         string
    Embedding       []float64 `gorm:"type:vector(768)"`
}

// insertPaper adds a new paper to the database with an embedding.
func insertPaper(paper *Paper) error {
	// Generate embedding for the paper content
	embedding, err := generateEmbedding(paper.Content)
	if err != nil {
			log.Printf("Failed to generate embedding for paper '%s': %v", paper.Title, err)
			return err
	}

	// Assign the generated embedding to the paper struct
	paper.Embedding = embedding

	// Insert the paper into the database
	if err := Db.Create(paper).Error; err != nil {
			log.Printf("Failed to insert paper '%s' into the database: %v", paper.Title, err)
			return err
	}

	log.Printf("Paper '%s' inserted successfully", paper.Title)
	return nil
}