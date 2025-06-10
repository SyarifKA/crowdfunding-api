package routers

import (
	"github.com/SyarifKA/crowdfunding-api/controllers"
	"github.com/SyarifKA/crowdfunding-api/middlewares"
	"github.com/gin-gonic/gin"
)

// func userRouter(r *gin.RouterGroup) {
// 	r.GET("", controllers.FindAllUsers)
// 	r.POST("", controllers.RegistUser)
// }

func userRouter(r *gin.RouterGroup) {
	r.GET("", middlewares.LogWithMessage("Find all users", controllers.FindAllUsers))
	r.POST("", middlewares.LogWithMessage("Register new user", controllers.RegistUser))
}
