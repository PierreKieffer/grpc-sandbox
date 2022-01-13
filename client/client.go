package main

import (
	"context"
	"github.com/PierreKieffer/grpc-sandbox/server"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	Route()
}

/*
Simple method
*/
func Client() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect on port 9000: %v", err)
	}

	defer conn.Close()

	s := server.NewEventServiceClient(conn)

	event := server.Event{
		Payload: "foobar",
	}
	resp, err := s.Ack(context.Background(), &event)
	if err != nil {
		log.Fatalf("gRPC call error: %v", err)
	}

	log.Printf("Reponse from server : %s", resp.Payload)
}

/*
Server side streaming
*/
func Subscribe() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect on port 9000: %v", err)
	}

	defer conn.Close()

	s := server.NewEventServiceClient(conn)

	event := server.Event{
		Payload: "foobar",
	}

	stream, err := s.Subscribe(context.Background(), &event)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Payload)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
}

/*
Client side streaming
*/
func Publish() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect on port 9000: %v", err)
	}

	defer conn.Close()

	s := server.NewEventServiceClient(conn)

	stream, err := s.Publish(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	data := []string{"foo", "bar"}

	for _, d := range data {
		event := server.Event{
			Payload: d,
		}
		err := stream.Send(&event)
		if err != nil {
			log.Fatalf("Stream send error %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("CloseAndRecv %v", err)
	}
	log.Printf("%v", reply)
}

/*
Bidirectional streaming
*/
func Route() {

	done := make(chan bool)

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect on port 9000: %v", err)
	}

	defer conn.Close()

	s := server.NewEventServiceClient(conn)

	stream, err := s.Route(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("Stream is closed")
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Payload)

		}
	}()

	data := []string{"foo", "bar"}

	for _, d := range data {
		event := server.Event{
			Payload: d,
		}
		err := stream.Send(&event)
		if err != nil {
			log.Fatalf("Stream send error %v", err)
		}
	}
	stream.CloseSend()
	<-done
	log.Println("finished")
}
