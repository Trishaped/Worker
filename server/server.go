package main

import (
	//"bufio"
	"fmt"
	//"io"
	//"log"
	"net"
	//"os"
	//"strings"
	"time"
)

type CoreClient struct {
	host   string
	in     chan string
	out    chan string
	socket net.Conn
	quit   chan bool
}

/*func (c *Client) Read(buffer []byte) bool {
    bytesRead, error := c.Conn.Read(buffer)
    if error != nil {
        c.Close()
        Log(error)
        return false
    }
    Log("Read ", bytesRead, " bytes")
    return true
}*/

func main() {

	fmt.Println("Trishaped worker")
	connect := Connect("localhost:7777")
	connected := <-connect
	println(connected)
}

func handleConnectError(timeout uint16, err error) bool {

	if err != nil {

		if neterr, ok := err.(net.Error); ok && (neterr.Timeout() || neterr.Temporary()) {
			println(neterr.Timeout())
			println(neterr.Temporary())
			println("Apply exponential backoff")
		}

		println(err.Error())
		return false
	}

	return true
}

func Connect(host string) chan bool {

	connect := make(chan bool)

	go func() {

		timeout := uint16(250)

		for {

			_, err := net.DialTimeout("tcp", host, time.Duration(timeout) * time.Millisecond)

			if !handleConnectError(err) {
				connect <- false
				return
			}

			connect <- true
		}

	}()

	return connect

	/*defer socket.Close()

	socket.SetReadDeadline(time.Now().Add(10 * time.Second))

	if _, err = socket.Write([]byte("Hello" + "\n")); err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Client input error: ", err)
			return true
		}

		line = strings.TrimRight(line, " \t\r\n")

		if _, err = socket.Write([]byte(line + "\n")); err != nil {

			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				println("Timeout")
			}

			println(err == io.EOF)
			panic(err)
		}

	}*/

}
