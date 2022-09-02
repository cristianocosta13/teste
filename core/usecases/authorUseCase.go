package usecases

import (
	"teste/core/domain"
	"teste/core/interfaces"
)

type authorUseCase struct {
	authorRepo interfaces.AuthorRepository
}

func NewAuthorUseCase(authorRepo interfaces.AuthorRepository) authorUseCase {
	return authorUseCase{
		authorRepo: authorRepo,
	}
}

func (t *authorUseCase) Get(id string) (*domain.Author, error) {
	author, err := t.authorRepo.Get(id)
	if err != nil {
		//log.Errorw("Error getting from repo", logging.KeyID, id, logging.KeyErr, err)
		return nil, err
	}

	return author, nil
}
