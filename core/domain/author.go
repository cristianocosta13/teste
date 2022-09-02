package domain

import "fmt"

type Author struct {
	ID          string
	Title       string
	Description string
}

func NewAuthor(id string, title, description string) *Author {
	return &Author{
		ID:          id,
		Title:       title,
		Description: description,
	}
}

func (a *Author) String() string {
	return fmt.Sprintf("%s - %s", a.Title, a.Description)
}
