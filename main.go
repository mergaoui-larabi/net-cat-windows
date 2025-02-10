package main

import (
	"fmt"
	"netcat/server"
)

func main() {
	server := server.NewServer(":8080")
	go server.Display()
	err := server.Start()
	if err != nil {
		fmt.Println("start fail to :", err)
	}
}
