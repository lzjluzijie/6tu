package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lzjluzijie/yitu/models"
)

func RegisterRouters(router *gin.Engine) {
	router.StaticFile("/", "./frontend/dist/index.html")
	router.StaticFile("/index.html", "./frontend/dist/index.html")
	router.Static("/js", "./frontend/dist/js")
	router.Static("/css", "./frontend/dist/css")

	router.GET("/t/:id/*type", GetTu)
	//TODO
	router.GET("/t/:id", GetTu)

	api := router.Group("/api")
	api.POST("/upload", Upload)
}

func GetTu(c *gin.Context) {
	id := c.Param("id")
	t := c.Param("type")

	tid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	tu, err := models.GetTuByID(tid)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if t == "/webp" || t == "webp" {
		c.Redirect(http.StatusMovedPermanently, tu.OneDriveWebpURL)
		return
	}

	c.Redirect(http.StatusMovedPermanently, tu.OneDriveURL)
	return
}
