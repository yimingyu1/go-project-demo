package list

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type UpdateTodoRequest struct {
	Id uint64 `json:"id"`
}

type DeleteTodoRequest struct {
	Id uint64 `json:"id"`
}
