package main

import (
	"fmt"
	"github.com/trishaped/worker/server"
)

func main() {
	fmt.Println("Trishaped worker")

	c := 2

	switch c {
	case 0:
		{
			fmt.Println("x")
		}
	case 1:
		{
			fmt.Println("y")
		}
	case 2:
		{
			server.Connect()
		}
	default:
		{
			fmt.Println("a")
		}
	}

}
