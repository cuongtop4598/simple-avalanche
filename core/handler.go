package core

import (
	"encoding/json"
	"fmt"

	"github.com/perlin-network/noise"
)

func HandleReceiveData(ctx noise.HandlerContext) error {
	if !ctx.IsRequest() {
		return nil
	}
	block := Block{}
	err := json.Unmarshal(ctx.Data(), &block)
	if err != nil {
		fmt.Println("invalid received data")
		return err
	}
	fmt.Println("received block index: ", block.Index)

	if block.IsValid(GetLastestBlock()) {
		DataChannel <- block
		ctx.Send([]byte("1"))
	} else {
		ctx.Send([]byte("0"))
	}
	return nil
}
