package main

import (
	"context"
	"fmt"
	p "github.com/mactsouk/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = ":8080"

func main() {
	conn, err := grpc.Dial(
		port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Dial: ", err)
		return
	}

	client := p.NewMessageServiceClient(conn)
	r, err := AboutToSayIt(context.Background(), client, "My Message!")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Response Text: ", r.Text)
	fmt.Println("Response SubText: ", r.Subtext)
}

func AboutToSayIt(ctx context.Context, m p.MessageServiceClient, text string) (*p.Response, error) {
	request := &p.Request{
		Text:    text,
		Subtext: "New message",
	}

	r, err := m.SayIt(ctx, request)
	if err != nil {
		return nil, err
	}

	return r, nil
}
