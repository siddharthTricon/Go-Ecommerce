package routes

import (
	"github.com/gon-gonic/gin"
	"github.com/siddharthTricon/go-ecommerce/controllers"
)

func UsdrRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/seacrh", controllers.SeacrhProductByQuery())
}
