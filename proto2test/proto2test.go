package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	protodata "StudentsTable/proto2test/protodata"

	"github.com/golang/protobuf/proto"
)

func main() {
	op := flag.String("op", "s", "s for server, c for client")
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto2Server()
	case "c":
		RunProto2Client()
	}
}

func RunProto2Server() {
	l, err := net.Listen("tcp", ":6767")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		go func(cm net.Conn) {
			defer cm.Close()
			data, err := ioutil.ReadAll(c)
			if err != nil {
				return
			}
			a := &protodata.Animal{}
			proto.Unmarshal(data, a)
			fmt.Printf("%+v", a)
		}(c)
	}
}

func RunProto2Client() {
	a := &protodata.Animal{
		Id:         proto.Int(1),
		AnimalType: proto.String("Raptor"),
		Nickname:   proto.String("raptor"),
		Zone:       proto.Int(3),
		Age:        proto.Int(20),
	}
	data, err := proto.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	SendData(data)
}

func SendData(data []byte) {
	c, err := net.Dial("tcp", "localhost:6767")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Write(data)
}
