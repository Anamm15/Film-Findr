package routes

import (
	"FilmFindr/controller"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func ReviewRoute(server *gin.Engine, reviewController controller.ReviewController, jwtService service.JWTService) {
	review := server.Group("/api/v1/reviews")
	{
		review.GET("/user/:id", reviewController.GetReviewByUserId)
		review.GET("/film/:id", reviewController.GetReviewByFilmId)
		review.POST("/", middleware.Authenticate(jwtService), reviewController.CreateReview)
		review.PUT("/:id", middleware.Authenticate(jwtService), reviewController.UpdateReview)
		review.PATCH("/:id/reaction", middleware.Authenticate(jwtService), reviewController.UpdateReaksiReview)
		review.DELETE("/:id", middleware.Authenticate(jwtService), reviewController.DeleteReview)
	}
}
