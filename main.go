package main

import (
	"avalanche/core"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/perlin-network/noise"
)

var PeersAddress []uint16

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	PeersAddress = []uint16{8087, 8081, 8082, 8083, 8084, 8085, 8086}
	core.Node, _ = noise.NewNode(noise.WithNodeBindPort(PeersAddress[port]))
	core.Node.Handle(core.HandleReceiveData)
	core.CheckError(core.Node.Listen())
	fmt.Println(core.Node.Addr())

	go func() {
		for {
			block := <-core.DataChannel
			go func() {
				if core.Decide(block, PeersAddress) {
					core.Blockchain = append(core.Blockchain, block)
					fmt.Println("lastest block index :", block.Index)
				}
			}()
			time.Sleep(5 * time.Second)
		}
	}()

	if port == 6 {
		block := core.NewBlock(core.GetLastestBlock(), 20)
		data, _ := json.Marshal(block)
		res, err := core.Node.Request(context.TODO(), ":"+strconv.Itoa(int(PeersAddress[0])), data)
		if err != nil {
			core.CheckError(err)
		}
		fmt.Printf("%s\n", res)
	}
	select {}
}
