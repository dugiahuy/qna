package models

// Question ..
type Question struct {
	ID      int    `json:"ID"`
	Content string `json:"Content"`
	// CreatedTime time.Time `json:"CreatedTime"`
}

type question []Question
