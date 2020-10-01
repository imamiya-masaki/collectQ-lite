package handler

import (
	"app/database"
	"app/models"

	"github.com/gin-gonic/gin"
)

func ApiGetAllQuetions(c *gin.Context) {
	posts := models.GetAllQuetion(database.GetDB())
	println("確認！！AA", posts)
	c.JSON(200, posts)
}

func ApiRegistQuetion(c *gin.Context) {
	req := &models.Quetion{}
	err := c.BindJSON(req)
	if err != nil {
		println(err)
	}
	println("req", req)
	req.CreateQuetion(database.GetDB())
	c.JSON(200, req)
}

func ApiGetQuetion(c *gin.Context) {
	posts := models.GetQuetion(database.GetDB())
	c.JSON(200, posts)
}
