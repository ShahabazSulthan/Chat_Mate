// /internal/interfaces/controllers/user_controller.go
package handler

import (
	"chatting/pkg/domain"
	"chatting/pkg/usecases/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserInteractor user.UserInteractor
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.UserInteractor.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (u *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := u.UserInteractor.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
