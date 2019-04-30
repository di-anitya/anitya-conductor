package grpcserver

import (
	"log"
	"net"

	pb "proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type service struct{}

func (s *service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("call from", req.Name)
	rsp := new(pb.HelloReply)
	rsp.Message = "Hello " + req.Name + "."
	return rsp, nil
}

func main() {

	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &service{})
	s.Serve(l)

}
