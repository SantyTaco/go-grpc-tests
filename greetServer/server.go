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
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

type server struct {}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("Greet function was invoqued with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName;

	opt := option.WithCredentialsFile("service-account-file.json")
	//config := &firebase.Config{ProjectID: "salespath"}
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	token, err := client.CustomToken(context.Background(), "some-uid")
	if err != nil {
		log.Fatalln( err)
	}
	fmt.Printf("Got custom token: %v\n", token)

	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func initializeFirebase() {
	opt := option.WithCredentialsFile("service-account-file.json")
	//config := &firebase.Config{ProjectID: "salespath"}
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	} else {
		fmt.Printf("Inizialize Success: ", app)
	}
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

	initializeFirebase()

	s := grpc.NewServer();
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}
