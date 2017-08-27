package main

import (
	"fmt"

	"github.com/kkishi/ticktacktoe/server"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

func main() {
	fmt.Printf("Server version: %s\n", server.Version)
	fmt.Printf("%v\n", &tpb.Cell{
		Row: 1,
		Col: 2,
	})
}
