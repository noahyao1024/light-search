package main

import (
	"github.com/gin-gonic/gin"
	"github.com/light-search/handler"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/:index/_search", handler.Search)
	r.PUT("/:index", handler.Index)
	r.POST("/:index", handler.Index)

	r.Run(":9200")
}
