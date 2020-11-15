package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmat-kurniawan/go-learn/controllers"
)

func Init() {

	r := gin.Default()

	movie := new(controllers.MovieController)
	r.GET("/movie", movie.GetAll)

	r.Run()
}
