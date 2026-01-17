package routes

import (
	"FilmFindr/controller"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func UserFilmRoutes(router *gin.Engine, userFilmController controller.UserFilmController, jwtService service.JWTService) {
	userFilms := router.Group("/api/v1/user-films")
	{
		userFilms.GET("/user/:id", userFilmController.GetUserFilmByUserId)
		userFilms.POST("/", middleware.Authenticate(jwtService), userFilmController.CreateUserFilm)
		userFilms.PATCH("/:id/status", middleware.Authenticate(jwtService), userFilmController.UpdateStatusUserFilm)
	}
}
