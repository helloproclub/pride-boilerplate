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
	Title       string `json:"title"`
	Description string `json:"description"`
}

type updateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type todoResponse struct {
	Ok   bool        `json:"ok"`
	Data domain.Todo `json:"data"`
}

type todoListResponse struct {
	Ok   bool          `json:"ok"`
	Data []domain.Todo `json:"data"`
}

func (handler *Handler) CreateNewTodo(ctx echo.Context) error {
	body := ctx.Request().Body
	var request createTodoRequest
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	todo, err := handler.service.Todo.CreateNewTodo(request.Title, request.Description)
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrInternalServerError
	}

	return ctx.JSON(200, todoResponse{
		Ok:   true,
		Data: todo,
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
		Ok:   true,
		Data: todos,
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
		Ok:   true,
		Data: todos,
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
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrInternalServerError
	}

	return ctx.JSON(200, todoResponse{
		Ok:   true,
		Data: todo,
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
		Ok:   true,
		Data: todo,
	})
}
