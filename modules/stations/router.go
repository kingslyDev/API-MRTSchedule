package stations

import "github.com/gin-gonic/gin"

func Initiate(router *gin.RouterGroup){
	Stations := router.Group("/stations")
	Stations.GET("", func(ctx *gin.Context) {
		// CODE SERVICES
	})
}