package handlers

import "teste/core/interfaces"

type TestHandler struct {
	authorUseCase interfaces.AuthorUseCase
}

func NewTodoHandle(authorUseCase interfaces.AuthorUseCase) *AuthorHandler {
	handler := &AuthorHandler{
		authorUseCase: authorUseCase,
	}
	return handler
}

func (ah *TestHandler) Get() int {
	//id := req.PathParameter("id")
	var id string
	_, err := ah.authorUseCase.Get(id)
	if err != nil {
		//resp.WriteError(500, err)
		return 8
	}

	//var todo *ToDo = &ToDo{}

	//todo.FromDomain(result)
	//resp.WriteAsJson(todo)
	return 8
}
