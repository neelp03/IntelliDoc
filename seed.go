package main

import (
    "log"
    "time"
)

func seedData() {
	papers := []Paper{
			{
					Title:           "Advancements in AI for Healthcare",
					Authors:         "Dr. Alice Smith, Dr. Bob Jones",
					PublicationDate: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
					Section:         "Abstract",
					Content:         "This paper discusses the latest advancements in AI applications within healthcare.",
			},
			{
					Title:           "Machine Learning Techniques for Medical Data",
					Authors:         "Dr. Carol Brown",
					PublicationDate: time.Date(2022, 5, 20, 0, 0, 0, 0, time.UTC),
					Section:         "Introduction",
					Content:         "Machine learning has revolutionized medical data analysis, providing faster and more accurate insights.",
			},
	}

	for i := range papers {
			if err := insertPaper(&papers[i]); err != nil {  // Pass a pointer to each paper
					log.Printf("Error inserting paper: %v", err)
			}
	}

	log.Println("Sample data seeded successfully.")
}

