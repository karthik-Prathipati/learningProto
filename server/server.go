package main

import (
	"context"
	"fmt"

	"github.com/karthik-Prathipati/learningProto/proto"
)

type server struct {
	proto.UnimplementedNumberManipulationServer
}

func (*server) Add(ctx context.Context, req *proto.Numbers) (resp *proto.Number, err error) {
	fmt.Println("Welcome to the Unary Add function")

	x := req.GetA()
	y := req.GetB()

	res := x + y
	resp = &proto.Number{
		A: res,
	}
	return resp, nil
}
