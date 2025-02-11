package domain

import (
	"time"
)

type Markdown struct {
	ArticleName string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time // able to be default value.
}

func (m Markdown) Read(p []byte) (n int, err error) {
	return copy(p, m.Content), nil
}

type NamedContent struct {
	Name    string
	Content []byte
}

type Article struct {
	ID        string
	StudentID string
	Title     string
	Content   string
	Imgs      []*NamedContent
	CreatedAt time.Time
	UpdatedAt time.Time
}
