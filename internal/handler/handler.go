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
			users.PUT("/:id/project/:projectId", h.deleteUserInProjectById)              //todo
			users.GET("/:id/projects", h.getProjectsByUserId)                            //todo
			users.POST("/:id/project/:projectId", h.createWorkerTaskInProject)           //todo
			users.GET("/:id/project/:projectId/tasks", h.getWorkersTasksInProject)       //todo
			users.GET("/:id/project/:projectId/activity", h.getWorkersActivityInProject) //todo
		}

		tasks := api.Group("/tasks")
		{
			tasks.GET("/:id/reports", h.getTaskReportsByTaskId)    //todo
			tasks.POST("/:id/reports", h.createTaskReportByTaskId) //todo
			tasks.GET("/:id", h.getTaskById)                       //todo
			tasks.PUT("/:id", h.closeTaskById)                     //todo
		}

		projects := api.Group("/projects")
		{
			projects.POST("/", h.createProject)
			projects.PUT("/:projectId/users/:id", h.addUserInProjectById) //todo
			projects.GET("/", h.getAllProjects)
			projects.GET("/:id", h.getProjectById)            // todo !!!returning closed projects too
			projects.GET("/:id/users", h.getUsersByProjectId) //todo
			projects.DELETE("/:id", h.deleteProject)
		}
	}

	return router
}
