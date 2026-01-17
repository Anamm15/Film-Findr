package routes

import (
	"FilmFindr/controller"
	"FilmFindr/helpers"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func FilmRoute(server *gin.Engine, filmController controller.FilmController, jwtService service.JWTService) {
	film := server.Group("/api/v1/films")
	{
		film.GET("/", filmController.GetAllFilm)
		film.GET("/:id", filmController.GetFilmById)
		film.GET("/get-top-film", filmController.GetTopFilm)
		film.GET("/get-trending-film", filmController.GetTrendingFilm)
		film.GET("/search", filmController.SearchFilm)
		film.POST("/", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.CreateFilm)
		film.PUT("/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.UpdateFilm)
		film.PATCH("/:id/status", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.UpdateStatus)
		film.DELETE("/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.DeleteFilm)
		film.POST("/add-film-genre", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.AddFilmGenre)
		film.DELETE("/delete-film-genre", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.DeleteFilmGenre)
	}
}
