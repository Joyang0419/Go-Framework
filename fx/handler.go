package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Demo1(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": h.service.Demo1(),
		},
	)
}

func (h *Handler) Demo2(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": h.service.Demo2(),
		},
	)
}
