package controllers

import (
	"example.com/letter-api/services"
	"github.com/gin-gonic/gin"
)

// controller will interact with servece so it will have user service reference
type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}
//will be acting as a handler or a route function

func (uc *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	 ctx.JSON(200 , "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	 ctx.JSON(200 , "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	 ctx.JSON(200 , "")
}
//RouterGroup is used internally to configure router, a RouterGroup is associated with a prefix and an array of handlers (middleware
func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PUT("/update", uc.UpdateUser)
	userroute.DELETE("/delete", uc.DeleteUser)
	 
}