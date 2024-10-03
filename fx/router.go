package main

import (
	"github.com/gin-gonic/gin"
)

type HandlerMapping struct {
	h *Handler
}

func NewHandlerMapping(
	h *Handler,
) *HandlerMapping {
	return &HandlerMapping{
		h: h,
	}
}

func SetupRouter(handlerMapping *HandlerMapping) *gin.Engine {
	router := gin.Default()

	router.GET("/demo1", handlerMapping.h.Demo1)
	router.GET("/demo2", handlerMapping.h.Demo2)

	return router
}
