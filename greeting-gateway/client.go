package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"

	pb "github.com/omar-khawaja/grpc-nomad-demo/hellopb"
)

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	name := r.FormValue("name")
	greeting := getGreeting(name)
	fmt.Fprintf(w, "%s\n", greeting)
}

func getGreeting(name string) string {
	address := os.Getenv("address")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	hr := &pb.HelloRequest{Name: name}
	r, err := c.Say(context.Background(), hr)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return r.Message
}

func main() {
	http.HandleFunc("/", greetingHandler)
	log.Println("Starting greeting gateway...")
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", nil))
}
