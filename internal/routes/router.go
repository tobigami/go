package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping/:id", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	id := c.Param("id")
	sort := c.Query("sort")
	filter := c.DefaultQuery("filter", "default_filter")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...pong...pong hehehe",
		"id":      id,
		"filter":  filter,
		"sort":    sort,
		"users":   []string{"thanh", "lam", "nhi"},
	})
}
