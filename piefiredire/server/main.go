package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"server/services"
)

func main() {
	var s *grpc.Server

	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse()

	if *tls {
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}
		s = grpc.NewServer(grpc.Creds(creds))
	} else {
		s = grpc.NewServer()
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterMeatCounterServer(s, services.NewMeatCounterServer())
	reflection.Register(s)

	fmt.Print("gRPC server listening on port 50051")
	if *tls {
		fmt.Println(" with TLS")
	} else {
		fmt.Println()
	}
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
