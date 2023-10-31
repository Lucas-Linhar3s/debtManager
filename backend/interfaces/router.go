package interfaces

import (
	"github.com/Lucas-Linhar3s/debtManager/interfaces/auth"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers(r *gin.RouterGroup) {
	// AUHT
	auth.Router(r.Group("/auth"))

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
