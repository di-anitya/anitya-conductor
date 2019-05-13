package main

import (
	"log"
	"net"

	pb "./proto"

	monitor "../monitoring"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type service struct{}

func (s *service) SendJobInfo(ctx context.Context, req *pb.JobRequest) (*pb.JobReply, error) {
	log.Println("exec..", req.Category, "-", req.TargetUrl)
	var status bool
	var result string

	if req.Category == "http" {
		status, result = monitor.RunHTTPValification(req.TargetUrl)
	} else if req.Category == "dns" {
		status, result = monitor.RunDNSValification(req.TargetUrl)
	} else {
		status = false
		result = "Unknown Job was defined." + " " + "(" + req.TargetUrl + ")"
	}

	rsp := new(pb.JobReply)
	rsp.Status = status
	rsp.Result = result
	return rsp, nil
}

func main() {

	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterMonitoringJobServer(s, &service{})
	s.Serve(l)

}
