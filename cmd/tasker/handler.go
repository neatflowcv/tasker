package main

import (
	"errors"
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
	Title       string `json:"title" binding:"required" example:"새로운 작업"`
	Description string `json:"description" example:"작업 설명"`
}

type TaskResponse struct {
	ID          string `json:"id" example:"1"`
	Title       string `json:"title" example:"새로운 작업"`
	Description string `json:"description" example:"작업 설명"`
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
func (h *Handler) CreateTask(ctx *gin.Context) {
	var req CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	spec := domain.NewTaskSpec(req.Title, req.Description)

	task, err := h.service.CreateTask(ctx, spec)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	response := TaskResponse{
		ID:          string(task.ID()),
		Title:       task.Title(),
		Description: task.Description(),
	}

	ctx.JSON(http.StatusCreated, response)
}

// ListTasks 모든 Task 목록 조회
// @Summary List tasks
// @Description 모든 Task 목록을 조회합니다
// @Tags tasks
// @Produce json
// @Success 200 {array} TaskResponse
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func (h *Handler) ListTasks(ctx *gin.Context) {
	tasks, err := h.service.ListTasks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	var responses []*TaskResponse
	for _, task := range tasks {
		responses = append(responses, &TaskResponse{
			ID:          string(task.ID()),
			Title:       task.Title(),
			Description: task.Description(),
		})
	}

	ctx.JSON(http.StatusOK, responses)
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
func (h *Handler) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID가 필요합니다"})

		return
	}

	task, err := h.service.GetTask(ctx, domain.TaskID(id))
	if err != nil {
		if errors.Is(err, core.ErrTaskNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task를 찾을 수 없습니다"})

			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	response := TaskResponse{
		ID:          string(task.ID()),
		Title:       task.Title(),
		Description: task.Description(),
	}

	ctx.JSON(http.StatusOK, response)
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
func (h *Handler) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID가 필요합니다"})

		return
	}

	var req CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	spec := domain.NewTaskSpec(req.Title, req.Description)

	ret, err := h.service.UpdateTask(ctx, domain.TaskID(id), spec)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	response := TaskResponse{
		ID:          string(ret.ID()),
		Title:       ret.Title(),
		Description: ret.Description(),
	}

	ctx.JSON(http.StatusOK, response)
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
func (h *Handler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID가 필요합니다"})

		return
	}

	err := h.service.DeleteTask(ctx, domain.TaskID(id))
	if err != nil {
		if errors.Is(err, core.ErrTaskNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task를 찾을 수 없습니다"})

			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.Status(http.StatusNoContent)
}
