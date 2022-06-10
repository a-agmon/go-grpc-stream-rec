package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/a-agmon/rcserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientAddr string = "localhost:9090"

func main() {

	conn, err := grpc.Dial(clientAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dialing %v", err)
	}
	defer conn.Close()
	grpcClient := pb.NewRcServiceClient(conn)
	runClient(grpcClient)

}

func runClient(client pb.RcServiceClient) {

	stream, err := client.Recommend(context.Background())
	if err != nil {
		log.Fatalf("cant open client stream %v \n", err)
	}
	
	// this is a routine that simple prints whatever it recieves from the server
	go func() {
		for {

			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("Recieved EOF")
			}
			if err != nil {
				log.Fatalf("error recieveing stream msg %v \n", err)
			}
			log.Println("==> we recieved: " + res.Content)
		}

	}()

	// now we block the main thread with using the main thread to wait for user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			stream.Send(&pb.RecommendRequest{
				Content: text,
			})
		} else {
			// exit if user entered an empty string
			break
		}
	}
	//start incomong stream loop

}
