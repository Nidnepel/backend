package handler

import (
	"github.com/Nidnepel/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id/project/:projectId", h.deleteUserInProjectById)
			users.GET("/:id/projects", h.getProjectsByUserId)
			users.POST("/:id/project/:projectId", h.createWorkerTaskInProject)
			users.GET("/:id/project/:projectId/tasks", h.getWorkersTasksInProject)
			users.GET("/:id/project/:projectId/activity", h.getWorkersActivityInProject)
			users.POST("/activity", h.sendWorkersActivityInProject)
		}

		tasks := api.Group("/tasks")
		{
			tasks.GET("/:id/reports", h.getTaskReportsByTaskId)
			tasks.POST("/:id/reports", h.createTaskReportByTaskId)
			tasks.GET("/:id", h.getTaskById)
			tasks.PUT("/:id", h.closeTaskById)
		}

		projects := api.Group("/projects")
		{
			projects.POST("/", h.createProject)
			projects.PUT("/:projectId/users/:id", h.addUserInProjectById)
			projects.GET("/", h.getAllProjects)
			projects.GET("/:id", h.getProjectById)
			projects.GET("/:id/users", h.getUsersByProjectId)
			projects.DELETE("/:id", h.deleteProject)
		}
	}

	return router
}
