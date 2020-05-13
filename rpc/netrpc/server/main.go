package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Listener int
type Reply struct {
	Data string
}

func (l *Listener) GetLine(line []byte, reply *Reply) error {
	rv := string(line)
	fmt.Println("%v\n", rv)

	*reply = Reply{rv}
	return nil
}
func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
