package handlers

import "teste/core/interfaces"

type AuthorHandler struct {
	authorUseCase interfaces.AuthorUseCase
}

func NewTodoHandler(authorUseCase interfaces.AuthorUseCase) *AuthorHandler {
	handler := &AuthorHandler{
		authorUseCase: authorUseCase,
	}
	return handler
}

func (ah *AuthorHandler) Get() {
	//id := req.PathParameter("id")
	var id string
	result, err := ah.authorUseCase.Get(id)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todo *ToDo = &ToDo{}

	todo.FromDomain(result)
	resp.WriteAsJson(todo)
}
