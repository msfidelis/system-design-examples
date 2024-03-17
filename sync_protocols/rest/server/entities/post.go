package entities

import "github.com/uptrace/bun"

type Post struct {
	bun.BaseModel `bun:"table:posts,alias:p"`
	ID            string `json:"id" bun:"id,pk,autoincrement"`
	Author        string `json:"autor"`
	Title         string `json:"title"`
	Content       string `json:"content"`
}
