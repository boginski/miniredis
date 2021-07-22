package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

/*
router.POST("/api/set/", handler.SetKeyValue)
router.GET("/api/get/", handler.GetKey)
router.GET("/api/get/pattern", handler.GetPatternKey)
router.DELETE("/api/delete/", handler.DeleteKey)
*/

func (h *Handler) SetKeyValue(c *gin.Context) {
	var kv KeyValue

	if err := c.BindJSON(&kv); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	err := h.storage.Set(&kv)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK"})
}

func (h *Handler) GetKey(c *gin.Context) {
	var rv RequestValue

	if err := c.BindJSON(&rv); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	answer, err := h.storage.Get(&rv)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, answer)
}

func (h *Handler) GetPatternKey(c *gin.Context) {
	var pv PatternValue

	if err := c.BindJSON(&pv); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	answer, err := h.storage.Keys(&pv)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, answer)
}

func (h *Handler) DeleteKey(c *gin.Context) {
	var rv RequestValue

	if err := c.BindJSON(&rv); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.storage.Delete(&rv)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK"})
}
