package main


import (
	"fmt"
	"net/http"
	"os/exec"
	"bytes"

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
	router := gin.Default()
	router.GET("/pinned", getPins)
	// router.Run("localhost:8080")
	router.Run("164.92.115.9:8081")
	fmt.Println("Started Router Module")
}

func getPins(c *gin.Context) {
	getPinsCmd := exec.Command("bash", "-c", "echo $PATH")
	var out bytes.Buffer
    getPinsCmd.Stdout = &out

	err := getPinsCmd.Run()

    if err != nil {
        fmt.Println(err)
    }

	fmt.Println(out.String())

	c.IndentedJSON(http.StatusOK, temp)
}

// func postSong()