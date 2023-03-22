package routers

import (
	"AuroraPixel/api"

	"github.com/gin-gonic/gin"
)

func (r Routers) imagesRouter(c *gin.RouterGroup) {
	ImagesApi := api.GroupApi.ImagesApi
	//单文件上传
	c.POST("images/upload", ImagesApi.Upload)
}
