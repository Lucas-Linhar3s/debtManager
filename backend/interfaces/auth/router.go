package auth

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/signup", SignUp)
	r.POST("/signin", SignIn)
	r.GET("/", Success)
}
