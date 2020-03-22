package main

import (
	"context"
	"fmt"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"log"
	"time"
)

func main() {
	const serverPort int = 7878
	// Connect to server
	RSocketConnect(serverPort)

}

func RSocketConnect( serverPort int){
	cli, err := rsocket.Connect().
		Resume().
		Fragment(1024).
		SetupPayload(payload.NewString("Hello", "World")).
		Transport(fmt.Sprintf("%s%d","tcp://127.0.0.1:", serverPort)).
		Start(context.Background())
	if err != nil {
		panic(err)
	}
	counter := 0
	var startTime = time.Now().Format(time.RFC3339)
	for i := 0; i<= 100000; i++ {
		// Send request
		result, err := cli.RequestResponse(payload.NewString(fmt.Sprintf("%s%d","count:",i), "Metadata")).
			Block(context.Background())
		if err != nil {
			panic(err)
		}
		log.Println("response:", result)
		counter += i
	}
	err = cli.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(startTime)
	fmt.Println(time.Now().Format(time.RFC3339))
}