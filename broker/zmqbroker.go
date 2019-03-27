//
//  Asynchronous REQ/REP broker
//

package broker

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"

	"log"
)

type backend interface {
	put()
}

type zmqb struct {
	uri string
}

func (z zmqb) put() {
	fmt.Println("trying to put message")
}

func newZmqb(uri string) *zmqb {
	z := new(zmqb)
	z.uri = uri
	return z
}

//  The following broker uses the multithreaded server model to deal requests out to a pool
//  of workers and route replies back to clients. One worker can handle
//  one request at a time but one client can talk to multiple workers at
//  once.

func CreateBroker(frontendURI string, backendType string, backendURI string) {

	//  Frontend socket talks to clients over TCP
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	frontend.Bind(frontendURI)

	// depending on the type of the backend, change who to send data to...

	//  Backend socket talks to workers over inproc
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	backend.Bind(backendURI)

	//  Connect backend to frontend via a proxy
	err := zmq.Proxy(frontend, backend, nil)
	log.Fatalln("Broker interrupted:", err)
}
