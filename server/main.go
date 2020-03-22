package main

import (
	"context"
	"fmt"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx/mono"
)

func main() {
	const echoPort int = 7878
	RSocketServer(echoPort)
}

func RSocketServer( port int ){
	println("Echo Server Started")
	err := rsocket.Receive().
		Resume().
		Fragment(1024).
		Acceptor(func(setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
			// bind responder
			return rsocket.NewAbstractSocket(
				rsocket.RequestResponse(func(msg payload.Payload) mono.Mono {
					println("Got a request")
					println(msg.DataUTF8())
					return mono.Just(msg)
				}),
			), nil
		}).
		Transport( fmt.Sprintf("%s%d","tcp://127.0.0.1:", port)).
		Serve(context.Background())
	panic(err)
}