package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cosmonaut/blog/x/blog/types"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

func main() {
	cosmos, err := cosmosclient.New(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	accountName := "alice"

	address, err := cosmos.Address(accountName)
	if err != nil {
		log.Fatal(err)
	}

	msg := &types.MsgCreatePost{
		Creator: address.String(),
		Title:   "Hello!",
		Body:    "This is the first post in here.",
	}

	txResp, err := cosmos.BroadcastTx(accountName, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("MsgCreatePost:\n\n")
	fmt.Println(txResp)

	queryClient := types.NewQueryClient(cosmos.Context)

	queryResp, err := queryClient.Posts(context.Background(), &types.QueryPostsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\n\nAll posts:\n\n")
	fmt.Println(queryResp)
}
