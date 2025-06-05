package routers

import "github.com/gin-gonic/gin"

func RoutersCombine(r *gin.Engine) {
	userRouter(r.Group("/users"))
}
