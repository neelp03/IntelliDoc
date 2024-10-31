package main

import (
	"log"
	"time"
)

type Paper struct {
	ID              int       `db:"id"`
	Title           string    `db:"title"`
	Authors         string    `db:"authors"`
	PublicationDate time.Time `db:"publication_date"`
	Section         string    `db:"section"`
	Content         string    `db:"content"`
	Embedding       []float64 `db:"embedding"`
}

// CreateSchema ensures the database schema exists
func createSchema() {
	schema := `
	CREATE TABLE IF NOT EXISTS papers (
		id SERIAL PRIMARY KEY,
		title TEXT,
		authors TEXT,
		publication_date DATE,
		section TEXT,
		content TEXT,
		embedding vector(768)  -- Adjust the dimension based on the model used
	);
	`
	_, err := Db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}
}

// InsertPaper adds a new paper to the database
func insertPaper(paper Paper) error {
	_, err := Db.Exec(`
		INSERT INTO papers (title, authors, publication_date, section, content, embedding)
		VALUES ($1, $2, $3, $4, $5, $6);
	`, paper.Title, paper.Authors, paper.PublicationDate, paper.Section, paper.Content, paper.Embedding)
	return err
}
