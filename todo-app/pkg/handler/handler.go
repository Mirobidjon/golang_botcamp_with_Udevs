package handler

import (
	"github.com/Mirobidjon/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
)

//Handler struct this
type Handler struct {
	service *service.Service
}

//constructur handler
func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// init routes with gin
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentify)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListByID)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.GET("/", h.getAllItems)
				items.POST("/", h.createItem)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemByID)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
