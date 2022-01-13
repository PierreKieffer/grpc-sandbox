package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

type Server struct {
	UnimplementedEventServiceServer
}

/*
Ack
Simple gRPC method
*/
func (s *Server) Ack(ctx context.Context, event *Event) (*Event, error) {
	log.Printf("Received event payload from client : %s", event.Payload)
	return &Event{Payload: "__ACK__"}, nil
}

/*
Subscribe
Server side streaming method
*/
func (s *Server) Subscribe(event *Event, stream EventService_SubscribeServer) error {
	for {
		resp := Event{Payload: "__ACK__"}
		err := stream.Send(&resp)
		if err != nil {
			log.Printf("send error %v", err)
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

/*
Subscribe
Client side streaming method
*/
func (s *Server) Publish(stream EventService_PublishServer) error {

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&Event{
				Payload: "End of stream",
			})
		}

		if err != nil {
			log.Printf("send error %v", err)
			return err
		}
		log.Printf("Received event payload from client : %s", data.Payload)
	}
}

/*
Route
Bidirectional streaming method
*/
func (s *Server) Route(stream EventService_RouteServer) error {

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return nil
		}

		log.Printf("Received event payload from client : %s", in.Payload)

		resp := Event{
			Payload: fmt.Sprintf("Message received : %v", in.Payload),
		}
		err = stream.Send(&resp)
		if err != nil {
			log.Printf("send error %v", err)
			return err
		}
	}
}
