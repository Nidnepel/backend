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

func (h *Handler) getProjectsByUserId(context *gin.Context) {

}

func (h *Handler) getWorkersTasksInProject(context *gin.Context) {

}

func (h *Handler) createWorkerTaskInProject(context *gin.Context) {

}

func (h *Handler) deleteUserInProjectById(context *gin.Context) {

}

func (h *Handler) getWorkersActivityInProject(context *gin.Context) {

}
