package task

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockService) GetByID(id int) (Task, error) {
	args := m.Called(id)
	return args.Get(0).(Task), args.Error(1)
}

func (m *MockService) Update(id int, task Task) (Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(Task), args.Error(1)
}

type Handler struct {
	service *MockService
}

func (h *Handler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid task ID"})
	}

	if err := h.service.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
	}

	return c.NoContent(http.StatusNoContent)
}

func TestDelete(t *testing.T) {
	e := echo.New()
	mockService := new(MockService)
	handler := &Handler{service: mockService}

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/tasks/invalid", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/tasks/:id")
		c.SetParamNames("id")
		c.SetParamValues("invalid")

		if assert.NoError(t, handler.Delete(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Invalid task ID")
		}
	})

	t.Run("Task Not Found", func(t *testing.T) {
		taskID := 1
		mockService.On("Delete", taskID).Return(echo.ErrNotFound)

		req := httptest.NewRequest(http.MethodDelete, "/tasks/"+strconv.Itoa(taskID), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/tasks/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(taskID))

		if assert.NoError(t, handler.Delete(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Contains(t, rec.Body.String(), "Not Found")
		}

		mockService.AssertExpectations(t)
	})
}
