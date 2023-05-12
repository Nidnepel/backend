package handler

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(c *gin.Context) {
	var input entity.User
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	id, err := h.services.User.CreateUser(ctx, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllUsers(c *gin.Context) {
	items, err := h.services.User.ReadAllUsers(context.Background())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result := make([]*entity.User, 0)
	for _, item := range items {
		if item.Status == entity.ActiveUser {
			result = append(result, item)
		}
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) getUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	item, err := h.services.User.ReadUser(ctx, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if item.Status == entity.NotActiveUser {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) getProjectsByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.services.User.ReadAllProjects(context.Background(), userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result := make([]*entity.Project, 0)
	for _, item := range items {
		if item.Status == entity.OpenedProject {
			result = append(result, item)
		}
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) getWorkersTasksInProject(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.services.User.GetTasksInProject(context.Background(), projectId, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result := make([]*entity.Task, 0)
	for _, item := range items {
		if item.ProgressStatus == entity.OpenedTask {
			result = append(result, item)
		}
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) createWorkerTaskInProject(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var newTask entity.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	id, err := h.services.User.CreateTask(ctx, projectId, userId, newTask)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteUserInProjectById(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	isOk, err := h.services.User.DeleteUserInProject(ctx, projectId, userId)
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

func (h *Handler) getWorkersActivityInProject(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	active, err := h.services.User.GetUserActivityInProject(ctx, projectId, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, active)
}

func (h *Handler) sendWorkersActivityInProject(c *gin.Context) {
	var newSession entity.Session
	if err := c.BindJSON(&newSession); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	id, err := h.services.User.CreateSession(ctx, newSession)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
