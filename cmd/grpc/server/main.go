package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strings"

	pr "github.com/sehogas/go_codes/cmd/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct {
	pr.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, req *pr.HelloRequest) (*pr.HelloResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty argument")
	}
	resp := &pr.HelloResponse{}
	resp.Language = req.Language
	resp.Message = translateHello(req.Language) + name + "!"
	return resp, nil
}

func (s *GreeterServer) SayHelloAgain(ctx context.Context, req *pr.HelloRequest) (*pr.HelloResponse, error) {
	return s.SayHello(ctx, req)
}

func (s *GreeterServer) SayReguards(ctx context.Context, req *pr.HelloRequest) (*pr.ListRegardsResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty argument")
	}
	resp := &pr.ListRegardsResponse{}
	resp.Items = append(resp.Items, &pr.HelloResponse{
		Language: pr.Language_ENGLISH,
		Message:  translateHello(pr.Language_ENGLISH) + name + "!",
	})
	resp.Items = append(resp.Items, &pr.HelloResponse{
		Language: pr.Language_SPANISH,
		Message:  translateHello(pr.Language_SPANISH) + name + "!",
	})
	return resp, nil
}

func translateHello(language pr.Language) string {
	switch language {
	case pr.Language_ENGLISH:
		return "Hello "
	case pr.Language_SPANISH:
		return "Hola "
	}
	return ""
}

func main() {

	log.SetPrefix("grpc server|")
	log.SetFlags(log.Ltime)

	port := flag.String("port", "8089", "Service listening port. Default 8089")
	flag.Parse()

	host := "0.0.0.0:" + *port

	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	grpcServer := grpc.NewServer()
	pr.RegisterGreeterServer(grpcServer, &GreeterServer{})

	log.Println("starting at", host)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("impossible to serve: %s\n", err)
	}
}
