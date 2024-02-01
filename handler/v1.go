package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/noahyao1024/light-gopkg/search"
)

func Search(c *gin.Context) {
	from, _ := strconv.ParseInt(c.DefaultQuery("from", "0"), 10, 64)
	size, _ := strconv.ParseInt(c.DefaultQuery("size", "10"), 10, 64)

	request := &search.V1Request{
		Query: analyzeQuery(c.DefaultQuery("q", "")),
		Index: c.Param("index"),
		From:  from,
		Size:  size,
	}

	response := search.V1(c, request)

	c.JSON(200, gin.H{
		"request":  request,
		"response": response,
		"message":  "pong",
	})
}

func Index(c *gin.Context) {
	index := c.Param("index")
	if index == "" {
		c.JSON(400, gin.H{
			"message": "index is required",
		})
		return
	}

	c.JSON(200, gin.H{
		"123":     search.V1Index(c, index),
		"456":     search.V1Index(c, index),
		"message": "pong",
	})
}

func Doc(c *gin.Context) {
	request := &search.V1Request{
		Index: c.Param("index"),
	}

	if request.Index == "" {
		c.JSON(400, gin.H{
			"message": "index is required",
		})
		return
	}

	if err := c.BindJSON(request); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	search.V1Put(c, request)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func analyzeQuery(q string) *search.V1RequestQuery {
	return &search.V1RequestQuery{
		Raw: q,
	}
}
