package interfaces

import "teste/core/domain"

// TodoUseCase assinaturas de métodos
type AuthorUseCase interface {
	Get(id string) (*domain.Author, error)
	List() ([]domain.Author, error)
	Create(title, description string) (*domain.Author, error)
}

// TodoRepository portas que exportam os repositorios
type AuthorRepository interface {
	Get(id string) (*domain.Author, error)
	List() ([]domain.Author, error)
	Create(todo *domain.Author) (*domain.Author, error)
}
