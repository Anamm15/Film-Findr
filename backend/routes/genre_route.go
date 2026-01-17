package routes

import (
	"FilmFindr/controller"
	"FilmFindr/helpers"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func GenreRoute(router *gin.Engine, genreController controller.GenreController, jwtService service.JWTService) {
	genre := router.Group("/api/v1/genres")
	{
		genre.GET("/", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), genreController.GetAllGenre)
		genre.POST("/", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), genreController.CreateGenre)
		genre.DELETE("/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), genreController.DeleteGenre)
	}
}
