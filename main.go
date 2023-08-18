package main

import (
	"context"
	"fmt"
	"log"

	"github.com/doffy007/golang-hashmap/pkg/server"
)

// @title 		Tag Golang Hashmap database API
// @version 	1.0
// @description A Tag Service for Golang Hashmap database API

// @host 		localhost:8081
// @BasePath 	/
func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run(ctx context.Context) error {
	server, err := server.NewServer()
	if err != nil {
		return err
	}
	err = server.Run(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return err
}
