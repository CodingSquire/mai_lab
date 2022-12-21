package handler

import (
	service "go_services_lab/pkg/user/service"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	services *service.ServiceUser
}

func NewHandlerUser(services *service.ServiceUser) *HandlerUser {
	return &HandlerUser{services: services}
}

func (h *HandlerUser) InitRoutesUser() *gin.Engine {
	router := gin.New()
	router.POST("/get/:id", h.getUserByID)
	router.POST("/del/:id", h.deleteUser)
	router.POST("/add", h.addUser)
	router.POST("/all", h.getUserList)

	return router
}
