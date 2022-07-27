package main

import (
	"fmt"
	"context"
	"net/http"
	"errors"

	ipfsClient "github.com/ipfs/go-ipfs-http-client"
	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/accesscontroller"
)



func InitDB() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbClient, err := ipfsClient.NewURLApiWithClient("localhost:5001", &http.Client{}) // uses ipfsClient package
	if err != nil {
		return err
	}

	db, err := orbitdb.NewOrbitDB(ctx, dbClient, nil)
	if err != nil {
		fmt.Printf("Failed to create orbitdb instance: %s\n", err)
		return err
	}

	options := orbitdb.CreateDBOptions{}
	ac := &accesscontroller.CreateAccessControllerOptions{Access: map[string][]string{"write": {"*"}}}
	options.AccessController = ac
	store, err := db.Docs(ctx, "kawa", &options)
	if err != nil {
		return err
	}

	fmt.Println(store.Address())
	for true {

	}
	return nil	
}

func ConnectToDB(address string, db orbitdb.OrbitDB, ctx context.Context) (orbitdb.DocumentStore, error) {
	options := orbitdb.CreateDBOptions{}
	ac := &accesscontroller.CreateAccessControllerOptions{Access: map[string][]string{"write": {"*"}}}
	options.AccessController = ac
	store, err := db.Docs(ctx, address, &options)
	if err != nil {
		return nil, err
	}

	if store.Address().String() != "/orbitdb/bafyreidslmv4pgy4x3d6rtwigc3jtdk3eggabagomuwxc7cqin4jb5ktjy/kawa" {
		return nil, errors.New("address doesn't match...")
	}
	return store, nil
}