package main

import (
	"context"
	"fmt"
	"log"

	"github.com/doffy007/golang-hashmap/pkg/server"
)

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
