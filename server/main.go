package main

import (
    pb "practice-grpc/server/go_grpc/proto"
)


func main() {
	pb.SayHello("ウィタス")
}