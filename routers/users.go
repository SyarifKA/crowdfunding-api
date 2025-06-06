package routers

import (
	"github.com/SyarifKA/crowdfunding-api/controllers"
	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.GET("", controllers.FindAllUsers)
	r.POST("", controllers.RegistUser)
}
