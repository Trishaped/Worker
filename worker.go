package main

import (
	"fmt"
	"github.com/trishaped/worker/server"
)

func main() {

	fmt.Println("Trishaped worker")
	server.Connect("localhost:7777")

}
