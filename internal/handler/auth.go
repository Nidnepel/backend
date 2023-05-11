package handler

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signIn(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	usr, err := h.services.Authorization.CheckUser(context.Background(), input.Login, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if usr.Status == entity.NotActiveUser {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error()+"abadlbmldf")
		return
	}

	c.JSON(http.StatusOK, usr)
}
