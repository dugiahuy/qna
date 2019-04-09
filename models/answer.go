package models

import "time"

// Answer ..
type Answer struct {
	Content     string
	QuestionID  int
	CreatedTime time.Time
}

// Array Answers ..
type answer []Answer
