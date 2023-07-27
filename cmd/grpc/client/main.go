package main

import (
	"context"
	"flag"
	"log"

	pr "github.com/sehogas/go_codes/cmd/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	log.SetPrefix("grpc client|")
	log.SetFlags(log.Ltime)

	port := flag.String("port", "8089", "Service listening port. Default 8089")
	flag.Parse()

	host := "0.0.0.0:" + *port

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect with server: %s\n", err.Error())
	}
	defer conn.Close()

	client := pr.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pr.HelloRequest{
		Name:     "Sebastian",
		Language: pr.Language_SPANISH,
	})
	if err != nil {
		log.Fatalf("error SayHello(): %s\n", err.Error())
	}
	log.Println("SayHello()", resp)

	resp, err = client.SayHelloAgain(context.Background(), &pr.HelloRequest{
		Name:     "Sebastian",
		Language: pr.Language_ENGLISH,
	})
	if err != nil {
		log.Fatalf("error SayHelloAgain(): %s\n", err.Error())
	}
	log.Println("SayHelloAgain()", resp)

	list, err := client.SayReguards(context.Background(), &pr.HelloRequest{
		Name:     "Sebastian",
		Language: pr.Language_SPANISH,
	})
	if err != nil {
		log.Fatalf("error SayReguards(): %s\n", err.Error())
	}
	log.Println("SayReguards()", list)

}
