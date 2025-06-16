package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neatflowcv/tasker/internal/app/flow"
	"github.com/neatflowcv/tasker/internal/pkg/domain"
	"github.com/neatflowcv/tasker/internal/pkg/repository/core"
)

type Handler struct {
	service *flow.Service
}

func NewHandler(service *flow.Service) *Handler {
	return &Handler{service: service}
}

type CreateTaskRequest struct {
	Title   string `json:"title" binding:"required" example:"새로운 작업"`
	Content string `json:"content" binding:"required" example:"작업 설명"`
}

type TaskResponse struct {
	ID      string `json:"id" example:"1"`
	Title   string `json:"title" example:"새로운 작업"`
	Content string `json:"content" example:"작업 설명"`
}

// CreateTask 새로운 Task 생성
// @Summary Create a new task
// @Description 새로운 Task를 생성합니다
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body CreateTaskRequest true "Task information"
// @Success 201 {object} TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func (h *Handler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	spec := domain.NewTaskSpec(req.Title, req.Content)
	task, err := h.service.CreateTask(c, spec)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := TaskResponse{
		ID:      string(task.ID()),
		Title:   task.Title(),
		Content: task.Content(),
	}

	c.JSON(http.StatusCreated, response)
}

// ListTasks 모든 Task 목록 조회
// @Summary List tasks
// @Description 모든 Task 목록을 조회합니다
// @Tags tasks
// @Produce json
// @Success 200 {array} TaskResponse
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func (h *Handler) ListTasks(c *gin.Context) {
	tasks, err := h.service.ListTasks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responses []TaskResponse
	for _, task := range tasks {
		responses = append(responses, TaskResponse{
			ID:      string(task.ID()),
			Title:   task.Title(),
			Content: task.Content(),
		})
	}

	c.JSON(http.StatusOK, responses)
}

// GetTask ID로 특정 Task 조회
// @Summary Get task
// @Description ID로 특정 Task를 조회합니다
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [get]
func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID가 필요합니다"})
		return
	}

	task, err := h.service.GetTask(c, domain.TaskID(id))
	if err != nil {
		if err == core.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task를 찾을 수 없습니다"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := TaskResponse{
		ID:      string(task.ID()),
		Title:   task.Title(),
		Content: task.Content(),
	}

	c.JSON(http.StatusOK, response)
}

// UpdateTask Task 수정
// @Summary Update task
// @Description Task를 수정합니다
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body CreateTaskRequest true "Updated task information"
// @Success 200 {object} TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [put]
func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID가 필요합니다"})
		return
	}

	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	spec := domain.NewTaskSpec(req.Title, req.Content)
	ret, err := h.service.UpdateTask(c, domain.TaskID(id), spec)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := TaskResponse{
		ID:      string(ret.ID()),
		Title:   ret.Title(),
		Content: ret.Content(),
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTask Task 삭제
// @Summary Delete task
// @Description Task를 삭제합니다
// @Tags tasks
// @Param id path string true "Task ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID가 필요합니다"})
		return
	}

	err := h.service.DeleteTask(c, domain.TaskID(id))
	if err != nil {
		if err == core.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task를 찾을 수 없습니다"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
