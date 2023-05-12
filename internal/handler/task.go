package handler

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getTaskReportsByTaskId(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.services.Task.ReadAllReports(context.Background(), taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) getTaskById(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	item, err := h.services.Task.ReadTask(ctx, taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if item.ProgressStatus == entity.ClosedTask {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) createTaskReportByTaskId(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var newTaskReport entity.TaskReport
	if err := c.BindJSON(&newTaskReport); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	id, err := h.services.Task.CreateReport(ctx, taskId, newTaskReport)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) closeTaskById(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	isOk, err := h.services.Task.CloseTask(ctx, taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if isOk == false {
		c.AbortWithStatusJSON(http.StatusBadRequest, isOk)
		return
	}

	c.JSON(http.StatusOK, isOk)
}
