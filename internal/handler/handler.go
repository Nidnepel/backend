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
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		managers := api.Group("/managers")
		{
			managers.POST("/", h.createManager)
			managers.GET("/", h.getAllManagers)
			managers.GET("/:id", h.getManagerById)
			managers.PUT("/:id/project/:projectId", h.deleteManagerInProjectById)
			managers.GET("/:id/projects", h.getProjectsByManagerId)

		}

		workers := api.Group("/workers")
		{
			workers.POST("/", h.createWorker)
			workers.POST("/:id/project/:projectId", h.createWorkerTaskInProject)
			workers.PUT("/:id/project/:projectId", h.deleteWorkerInProjectById)
			workers.GET("/", h.getAllWorkers)
			workers.GET("/:id", h.getWorkerById)
			workers.GET("/:id/projects", h.getProjectsByWorkerId)
			workers.GET("/:id/project/:projectId/tasks", h.getWorkersTasksInProject)
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
			projects.PUT("/:projectId/worker/:id", h.addWorkerInProjectById)
			projects.PUT("/:projectId/manager/:id", h.addManagerInProjectById)
			projects.GET("/", h.getAllProjects)
			projects.GET("/:id", h.getProjectById)
			projects.GET("/:id/workers", h.getWorkersByProjectId)
			projects.GET("/:id/managers", h.getManagersByProjectId)
			projects.DELETE("/:id", h.deleteProject)
		}
	}

	return router
}
