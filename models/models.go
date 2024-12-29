package models

import "time"

type Recipe struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Country     string    `json:"country"`
	Categories  string    `json:"categories"` // Comma-separated values
	CreatedAt   time.Time `json:"created_at"`
	Duration    string    `json:"duration"`
	Difficulty  string    `json:"difficulty"`
	Images      string    `json:"images"` // Comma-separated image URLs
	RecipeSteps string    `json:"recipe"` // Recipe steps as a string
}
