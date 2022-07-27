package main


import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Temp struct {
	ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Album  string  `json:"album"`
}

var temp = []Temp{
	{ID: "1", Title: "Go", Artist: "Flume", Album: "Palaces"},
}

func main() {
	fmt.Println("Starting Router Module")

	router := gin.Default()
	router.GET("/pinned", getPins)

	router.Run("localhost:8080")
}

func getPins(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, temp)
}