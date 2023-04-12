package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jdramirezl/proyecto-1-topicos/mom/pkg/client"
)

func main() {

	hostIP := os.Getenv("MOM_HOST")
	hostPort := os.Getenv("MOM_PORT")
	// time.Sleep(40 * time.Second)
	selfIP := os.Getenv("SELF_IP")
	momClient := client.NewClient(hostIP, hostPort, selfIP)
	momClient.AddConnection(selfIP)
	momClient.SubscribeQueue("default", selfIP)
	momClient.Connect("default")
	for {
		msg, err := momClient.ReceiveQueueMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	}
}
