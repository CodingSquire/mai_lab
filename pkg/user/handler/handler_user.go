package handler

import (
	"go_services_lab/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserID struct {
	ID int `uri:"id"`
}

func (h *HandlerUser) getUserByID(c *gin.Context) {
	var input UserID

	if err := c.ShouldBindUri(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user, err := h.services.User.Get(input.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":       user.ID,
			"name":     user.Name,
			"login":    user.Login,
			"password": user.Password,
		})
	}
}

func (h *HandlerUser) deleteUser(c *gin.Context) {
	var input UserID

	if err := c.ShouldBindUri(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.User.Delete(input.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

func (h *HandlerUser) addUser(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.User.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *HandlerUser) getUserList(c *gin.Context) {
	userList, err := h.services.User.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for i := range userList {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":    userList[i].ID,
			"name":  userList[i].Name,
			"login": userList[i].Login,
		})
	}
}
