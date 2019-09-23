package main

import (
	"flag"
	"log"
	"time"

	pb "github.com/marceloaguero/emojize-proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var server = flag.String("server", "localhost:9000", "The server's address")
var text = flag.String("text", "Hello world!", "The input text")

func init() {
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect to the service: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := pb.NewEmojiServiceClient(conn)

	log.Printf("Request: %s", *text)
	res, err := c.InsertEmojis(ctx, &pb.EmojiRequest{
		InputText: *text,
	})
	if err != nil {
		log.Fatalf("Couldn't call service: %v", err)
	}
	log.Printf("Server says: %s", res.OutputText)
}
