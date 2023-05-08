package handler

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createWorker(c *gin.Context) {
	var input entity.User
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	id, err := h.services.Worker.CreateWorker(ctx, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllWorkers(context *gin.Context) {

}

func (h *Handler) getWorkerById(c *gin.Context) {
	workerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	item, err := h.services.Worker.ReadWorker(ctx, workerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) getProjectsByWorkerId(context *gin.Context) {

}

func (h *Handler) getWorkersTasksInProject(context *gin.Context) {

}

func (h *Handler) createWorkerTaskInProject(context *gin.Context) {

}

func (h *Handler) deleteWorkerInProjectById(context *gin.Context) {

}
