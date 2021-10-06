package models

type CreateAttachmentModel struct {
	IdTodolist int    `json:"id_todolist"`
	Url        string `json:"url"`
	Caption    string `json:"caption"`
}

type UpdateAttachmentModel struct {
	Url     string `json:"url"`
	Caption string `json:"caption"`
}