package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/helloproclub/pride-boilerplate/domain"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type createTodoRequest struct {
	title       string
	description string
}

type updateTodoRequest struct {
	title       string
	description string
}

type todoResponse struct {
	ok   bool
	data domain.Todo
}

type todoListResponse struct {
	ok   bool
	data []domain.Todo
}

func (handler *Handler) CreateNewTodo(ctx echo.Context) error {
	body := ctx.Request().Body
	var request createTodoRequest
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	todo, err := handler.service.Todo.CreateNewTodo(request.title, request.description)
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrInternalServerError
	}

	return ctx.JSON(200, todoResponse{
		ok:   true,
		data: todo,
	})
}

func (handler *Handler) FindAllTodo(ctx echo.Context) error {
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	offset, err := strconv.Atoi(ctx.QueryParam("offset"))
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	todos, err := handler.service.Todo.FindAll(limit, offset)
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrInternalServerError
	}

	return ctx.JSON(200, todoListResponse{
		ok:   true,
		data: todos,
	})
}

func (handler *Handler) GetTodoByID(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	todos, err := handler.service.Todo.GetById(id)
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrInternalServerError
	}

	return ctx.JSON(200, todoResponse{
		ok:   true,
		data: todos,
	})
}

func (handler *Handler) UpdateTodoByID(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	body := ctx.Request().Body
	var request updateTodoRequest
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	todo, err := handler.service.Todo.Update(domain.Todo{
		Id:          id,
		Title:       request.title,
		Description: request.description,
	})
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrInternalServerError
	}

	return ctx.JSON(200, todoResponse{
		ok:   true,
		data: todo,
	})
}

func (handler *Handler) DeleteTodoByID(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	todo, err := handler.service.Todo.DeleteByID(id)
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrInternalServerError
	}

	return ctx.JSON(200, todoResponse{
		ok:   true,
		data: todo,
	})
}
