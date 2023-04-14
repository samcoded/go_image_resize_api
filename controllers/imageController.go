package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/samcoded/go_image_resize_api/utils"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to image resize api",
	})
}

func ImagesResize(c *gin.Context) {
	url := c.Query("url")
	width := c.Query("w")
	height := c.Query("h")

	resiveImg, err := utils.Resize(url, width, height)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// serve the image to the client
	c.Header("Content-Type", "image/jpeg")
	c.File(resiveImg)

	// c.JSON(200, gin.H{
	// 	"url":    url,
	// 	"width":  width,
	// 	"height": height,
	// 	"img":    "dd",
	// })
}
