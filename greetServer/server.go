package main

import (
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
	"../greetpb"
	"context"
	//"strconv"
	//"time"
	//"strconv"
	//"time"
)

type server struct {}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("Greet function was invoqued with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName;
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

/*func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error{
	fmt.Printf("GreetManyTimes function was invoqued with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName +" number" + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)//Send the response from the server
		time.Sleep(1000 * time.Millisecond)
	}
	return nil //Return nil when finish to send the response
}*/


func main() {
	fmt.Print("Hi, I am server!!")
	lis,err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil{
		log.Fatalf("Failed Listener: %v", err)
	}

	s := grpc.NewServer();
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}
