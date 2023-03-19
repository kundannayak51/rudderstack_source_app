package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rudderstack_source_app/entity"
	"github.com/rudderstack_source_app/services/user"
	"github.com/rudderstack_source_app/utils"
	"net/http"
)

type UserController struct {
	userService user.UserServiceInterface
}

func NewUserController(userService user.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (con *UserController) CreateOrUpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := utils.GetValueOnlyRequestContext(c)

	err := con.userService.CreateOrUpdateUser(ctx, user)

	if err != nil {

	}

	c.JSON(http.StatusOK, user)

}
