package handler

import "github.com/helloproclub/pride-boilerplate/service"

type Handler struct {
	service service.Service
}

func Init(service service.Service) Handler {
	return Handler{
		service: service,
	}
}
