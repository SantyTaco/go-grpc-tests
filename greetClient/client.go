package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"../greetpb"
	"context"
	//"io"
	//"io"
)

func main() {
	fmt.Print("Hi I am client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Could not connect: %v", err)
	}

	c := greetpb.NewGreetServiceClient(cc);
	fmt.Print("Created client grpc: %f", c)

	doUnary(c)
	//doServerStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient)  {
	fmt.Printf("Starting to do Unary GRPc")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Erin",
			LastName: "Moore",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet rps: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

/*func doServerStreaming(c greetpb.GreetServiceClient){
	fmt.Printf("Starting to do a server streaming rpc ")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Erin",
			LastName: "Moore",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes rps: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if (err == io.EOF){
			// we've reached the end of the stream
			break
		}
		if (err != nil) {
			log.Fatalf("Error while reading stream: %v",err)
		}
		log.Printf("Response from GreetManyTimes: &v", msg.GetResult())
	}
}*/

