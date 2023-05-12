package handler

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createProject(c *gin.Context) {
	var input entity.Project
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	id, err := h.services.Project.CreateProject(ctx, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllProjects(c *gin.Context) {
	items, err := h.services.Project.ReadAllProjects(context.Background())
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

func (h *Handler) getProjectById(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	item, err := h.services.Project.ReadProject(ctx, projectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) deleteProject(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	isOk, err := h.services.Project.CloseProject(ctx, projectId)
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

func (h *Handler) getUsersByProjectId(context *gin.Context) {

}

func (h *Handler) addUserInProjectById(context *gin.Context) {

}
