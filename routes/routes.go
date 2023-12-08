package routes

import (
	"awesomeProject/api"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", CreateUserHandler)
	r.PUT("/users/generateotp", GenerateOTPHandler)
	r.POST("/users/verifyotp", VerifyOTPHandler)
}

func CreateUserHandler(c *gin.Context) {
	apiHandler := c.MustGet("apiHandler").(*api.Handler)
	apiHandler.CreateNewUser()
}

func GenerateOTPHandler(c *gin.Context) {
	apiHandler := c.MustGet("apiHandler").(*api.Handler)
	apiHandler.GenerateOTP()
}

func VerifyOTPHandler(c *gin.Context) {
	apiHandler := c.MustGet("apiHandler").(*api.Handler)
	apiHandler.VerifyOTP()
}
