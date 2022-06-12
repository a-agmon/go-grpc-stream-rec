package main

import (
	"context"
	"io"
	"log"
	"strings"

	vm "github.com/a-agmon/rcserver/model"
	pb "github.com/a-agmon/rcserver/proto"
)

type RCServer struct {
	pb.RcServiceServer
	Handler *vm.EmbeddingHandler
}

func NewRCServer() *RCServer {

	vec_file := "/Users/alonagmon/MyData/work/golang-projects/vectors_model/factors.csv"
	item_file := "/Users/alonagmon/MyData/work/golang-projects/vectors_model/artists.csv"
	vec_size := 129

	return &RCServer{
		Handler: vm.NewEmbeddingHandler(vec_file, item_file, vec_size),
	}
}

func (s *RCServer) RecommendByMany(ctx context.Context, req *pb.RecommendManyRequest) (*pb.RecommendManyResponse, error) {

	maxProgs := 5
	viewed_items := req.Content
	if len(viewed_items) > maxProgs {
		viewed_items = viewed_items[len(viewed_items)-maxProgs:]
	}
	recommended, itemsNotfound, err := s.Handler.Recommend(viewed_items)
	if err != nil {
		return nil, err
	}
	log.Printf("Recieved request for: %v \n Reccomended:%v \n Not Found: %v \n -----\n", viewed_items, recommended, itemsNotfound)
	return &pb.RecommendManyResponse{
		Content: recommended,
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
