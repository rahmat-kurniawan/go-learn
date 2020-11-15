package controllers

import "github.com/gin-gonic/gin"

type MovieController struct{}

func (ctrl MovieController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{"page": "movie show all", "data": "data"})
	return
}
