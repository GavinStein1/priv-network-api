package main


import (
	"fmt"
	"net/http"
	"os/exec"
	"bytes"
	"context"

	"github.com/gin-gonic/gin"
	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/accesscontroller"
	// "berty.tech/go-orbit-db/iface"
	ipfsClient "github.com/ipfs/go-ipfs-http-client"
)

type Song struct {
	ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Album  string  `json:"album"`
}

type StdOut struct {
	Out    string  `json:"out"`
}

var temp = []Song{
	{ID: "1", Title: "Go", Artist: "Flume", Album: "Palaces"},
}

func checkErr(err error) bool {
	if err != nil {
		print(err)
		return true
	}
	return false
}

func CreateIPFSNode() (*ipfsClient.HttpApi, error) {
	// Create/connect to an IPFS node
	client, err := ipfsClient.NewURLApiWithClient("localhost:5001", &http.Client{}) // uses ipfsClient package
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CreateDBInstance(ctx context.Context, client *ipfsClient.HttpApi) (orbitdb.OrbitDB, error) {
	// Create an instance of orbitdb
	db, err := orbitdb.NewOrbitDB(ctx, client, nil)
	if err != nil {
		fmt.Printf("Failed to create orbitdb instance: %s\n", err)
		return nil, err
	}
	return db, nil
}

func ConnectToDocStore(ctx context.Context, db orbitdb.OrbitDB, address string) (orbitdb.DocumentStore, error) {
	// Connect to Kawa orbit document store
	options := orbitdb.CreateDBOptions{}
	ac := &accesscontroller.CreateAccessControllerOptions{Access: map[string][]string{"write": {"*"}}}
	options.AccessController = ac
	store, err := db.Docs(ctx, address, &options)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func getAll(doc interface{}) (bool, error) {
	return true, nil
}

func main() {
	err := main.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	// client, err := CreateIPFSNode()
	// if checkErr(err) {
	// 	return
	// }

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// // var db *orbitdb.OrbitDB
	// db, err := CreateDBInstance(ctx, client)
	// if checkErr(err) {
	// 	return
	// }

	// store, err := ConnectToDocStore(ctx, db, "/orbitdb/bafyreiajnlkjrqxhyxjnzg3wljvskqvpnrvhivztiyoonqinm2i2kxsdv4/pieces")
	// if checkErr(err) {
	// 	return
	// }
	
	// doc := temp[0]
	// var docMap map[string]interface{}
	// docMap = make(map[string]interface{})
	
	// docMap["_id"] = doc.ID
	// docMap["title"] = doc.Title
	// docMap["artist"] = doc.Artist
	// docMap["album"] = doc.Album

	// op, err := store.Put(ctx, docMap)
	// if err != nil {
	// 	return
	// }

	// fmt.Println(op.GetEntry())

	// cid := op.GetEntry().GetHash()
	// client.Pin()

	// results, err := store.Query(ctx, getAll)
	// if checkErr(err) {
	// 	return
	// }
	// fmt.Println(results)



	router := gin.Default()
	router.GET("/pinned", getPins)
	// router.Run("localhost:8080")
	router.Run("164.92.115.9:8081")
	fmt.Println("Started Router Module")
}

func ExecBashCommand(cmdString string) (bytes.Buffer, error) {
	cmd := exec.Command("bash", "-c", cmdString)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return bytes.Buffer{}, err
	}
	return out, nil
}

func getPins(c *gin.Context) {
	out, err := ExecBashCommand("ipfs pin ls")
	if err != nil {
        return
    }
	stdout := StdOut{out.String()}

	c.IndentedJSON(http.StatusOK, stdout)
}

// func postSong()