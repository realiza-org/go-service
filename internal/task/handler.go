package task

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service Service
}

func NewTaskHandler(service Service) *TaskHandler {
	return &TaskHandler{service: service}
}

func RegisterRoutes(g *echo.Group) {
	repo := NewInMemoryRepository()
	service := NewService(repo)
	handler := NewTaskHandler(service)

	g.GET("", handler.GetAll)
	g.POST("", handler.Create)
	g.GET("/:id", handler.GetByID)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}

func (h *TaskHandler) GetAll(c echo.Context) error {
	tasks := h.service.GetAll()
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) Create(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	createdTask := h.service.Create(task)
	return c.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
	}

	task, err := h.service.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
	}

	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updatedTask, err := h.service.Update(id, task)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
	}

	if err := h.service.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
