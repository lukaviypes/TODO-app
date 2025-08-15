package main

import (
	"awesomeProject/internal/api"
	"log"
)

func main() {
	addr := "3030"
	server := api.NewServer(addr)
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}

}
