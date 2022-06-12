package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	pb "github.com/a-agmon/rcserver/proto"
)

type RCServer struct {
	pb.RcServiceServer
}

func (s *RCServer) RecommendByMany(ctx context.Context, req *pb.RecommendManyRequest) (*pb.RecommendResponse, error) {

	viewed_items := req.Content
	req_str := fmt.Sprintf("RecommendByMany called with %#v", viewed_items)
	return &pb.RecommendResponse{
		Content: req_str,
	}, nil

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
