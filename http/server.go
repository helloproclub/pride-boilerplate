package http

import (
	"fmt"

	"github.com/helloproclub/pride-boilerplate/http/handler"
	"github.com/helloproclub/pride-boilerplate/service"
	"github.com/labstack/echo/v4"
)

type HTTPServer struct {
	server  *echo.Echo
	handler handler.Handler
}

/// Init is use to create a new server instance
func Init(service service.Service) HTTPServer {
	return HTTPServer{
		server:  echo.New(),
		handler: handler.Init(service),
	}
}

func (http *HTTPServer) Serve() {
	todo := http.server.Group("/todo")
	todo.POST("/", http.handler.CreateNewTodo)
	todo.GET("/", http.handler.FindAllTodo)
	todo.GET("/:id", http.handler.GetTodoByID)
	todo.PUT("/:id", http.handler.UpdateTodoByID)
	todo.DELETE("/:id", http.handler.DeleteTodoByID)

	fmt.Println(http.server.Start(":8000"))
}
