package main

import (
	"io"
	"log"
	"strings"

	pb "github.com/a-agmon/rcserver/proto"
)

type RCServer struct {
	pb.RcServiceServer
}

func (s *RCServer) Recommend(stream pb.RcService_RecommendServer) error {

	log.Println("Recommend called")

	var buffer_size = 3
	var buffer []string = make([]string, 0)
	for {
		in_stream, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		content := in_stream.Content
		log.Printf("Received: %s", content)
		buffer = append(buffer, content)
		log.Printf("current buffer: %v", buffer)
		if len(buffer) > buffer_size {
			buffer = buffer[1:]
		}
		log.Printf("new buffer: %v", buffer)
		err = stream.Send(&pb.RecommendResponse{
			Content: strings.Join(buffer, " "),
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
