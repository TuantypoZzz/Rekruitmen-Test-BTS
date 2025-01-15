package model

type ChecklistModel struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Username string `json:"username"`
}

type CreateChecklist struct {
	Title       string `json:"title" validate:"required,max=32"`
	Description string `json:"description" validate:"required"`
}
