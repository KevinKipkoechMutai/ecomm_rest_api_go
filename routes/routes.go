package routes

import (
	"github.com/akhil/ecommerce-yt/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup", controllers.signup)
	incomingRoutes.POST("/users/login", controllers.login)
	incomingRoutes.POST("/admin/addproduct", controllers.addproduct)
	incomingRoutes.GET("/users/productview", controllers.productview)
	incomingRoutes.GET("/users/search", controllers.search)
}


