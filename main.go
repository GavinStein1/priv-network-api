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

type StdOut struct {
	Out    string  `json:"out"`
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
	getPinsCmd := exec.Command("bash", "-c", "ipfs pin ls")
	var out bytes.Buffer
    getPinsCmd.Stdout = &out

	err := getPinsCmd.Run()

	stdout := StdOut{out.String()}

    if err != nil {
        fmt.Println(err)
    }

	c.IndentedJSON(http.StatusOK, stdout)
}

// func postSong()