package dtos

type NewPost struct {
	Title   string `json:"title" validate:"required"`
	Author  string `json:"author" validate:"required"`
	Content string `json:"content" validate:"required"`
}
