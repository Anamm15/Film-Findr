package routes

import (
	"FilmFindr/controller"
	"FilmFindr/helpers"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(server *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	user := server.Group("/api/v1/users")
	{
		user.GET("/", userController.GetAllUser)
		user.GET("", userController.GetUserByUsername)
		user.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		user.POST("/", userController.RegisterUser)
		user.POST("/login", userController.LoginUser)
		user.POST(("/logout"), middleware.Authenticate(jwtService), userController.LogoutUser)
		user.PATCH("/:id", middleware.Authenticate(jwtService), userController.UpdateUser)
		user.DELETE("/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), userController.DeleteUser)
	}
}
