package main

import (
	"context"
	"fmt"
	"io"
	"math"

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
func (*server) PrimeNumbers(req *proto.Number, resp proto.NumberManipulation_PrimeNumbersServer) error {
	fmt.Println("Welcome to the Prime numbers server")
	n := req.GetA()
	nums := make([]bool, n+1, n+1)
	for x := range nums {
		nums[x] = true
	}

	var i, j int64
	for i = 2; i <= n; i++ {
		if nums[i] {
			for j = i * i; j <= n; j += i {
				nums[j] = false
			}
		}
	}
	for i = 1; i < n+1; i++ {
		if nums[i] {
			res := proto.Number{
				A: i,
			}

			resp.Send(&res)
		}
	}
	return nil

}
func (*server) ComputeAverage(resp proto.NumberManipulation_ComputeAverageServer) error {
	fmt.Println(" Welcome to the Prime numbers server ")
	var avg, count int64 = 0, 0

	for {
		msg, err := resp.Recv()

		if err == io.EOF {
			return resp.SendAndClose(&proto.Number{
				A: avg,
			})
		}
		if err != nil {
			return fmt.Errorf("error at compute average func stream")
		}
		inp := msg.GetA()
		avg = avg*count + inp
		count++
		avg /= count
	}

}

func (*server) FindMaxNumber(resp proto.NumberManipulation_FindMaxNumberServer) error {
	fmt.Println("Currently in the FindMax function server")
	var max int64 = math.MinInt64

	for {
		req, err := resp.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Errorf("error at the receive func of the findMax stream")
			return err
		}
		val := req.GetA()
		if max < val {
			max = val

			err := resp.Send(&proto.Number{
				A: max,
			})

			if err != nil {
				fmt.Println("Error at the send func of the FIndMaxNumber")
				return err
			}
		}

	}
}
