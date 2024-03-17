package controllers

import (
	"fmt"
	"net/http"

	"main/dtos"

	"main/entities"

	"main/services"

	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {

	p := new(dtos.NewPost)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println(p)

	post := &entities.Post{
		Author:  p.Author,
		Title:   p.Title,
		Content: p.Content,
	}

	_, err := services.CreatePost(*post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "Hello, World!")
}
