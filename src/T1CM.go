package T1ConnectionManager

import (
	"fmt"
	"time"
)

func Run(client *T1Client, server *T1Server) {
	go runClient(client)
	go runServer(server)
}

func runClient(client *T1Client) {
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("I'm the client manager")
	}
}

func runServer(server *T1Server) {
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("I'm the server manager")
	}
}
