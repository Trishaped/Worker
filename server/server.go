package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Connect() {

	socket, err := net.Dial("tcp", "localhost"+":8080")

	if err != nil {
		panic(err)
	}

	defer socket.Close()

	if _, err = socket.Write([]byte("Hello"+"\n")); err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		line, err := reader.ReadString('\n')

		if err != nil {
            fmt.Println(err)
			return
		}

		line = strings.TrimRight(line, " \t\r\n")

        if _, err = socket.Write([]byte(line+"\n")); err != nil {
            panic(err)
        }

	}

}
