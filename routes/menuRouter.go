package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRouter(incomingRoute *gin.Engine) {
	incomingRoute.GET("/menus", controller.GetMenus())
	incomingRoute.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoute.POST("/menus", controller.CreateMenu())
	incomingRoute.PUT("/menus/:menu_id", controller.UpdateMenu())
}
