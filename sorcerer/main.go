package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//"sorcerer"
)

func doSomeWork() {

	threads := os.Getenv("SORCERER_THREADS")
	if threads == "" {
		threads = "1"
	}
	timeout := os.Getenv("SORCERER_TIMEOUT")
	if timeout == "" {
		timeout = "2"
	}

	cmd := exec.Command("stress", "--cpu", threads, "-v", "--timeout", timeout)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("sorcerer: exec error ", err)
	}
}

func generateHeat(w http.ResponseWriter, r *http.Request) {

	remote := r.RemoteAddr
	if r.Header.Get("X-FORWARDED-FOR") != "" {
		remote = r.Header.Get("X-FORWARDED-FOR")
	}

	log.Printf("sorcerer: handling request %s from %s", r.URL.Path, remote)
	doSomeWork()
	log.Printf("sorcerer: done with request %s from %s", r.URL.Path, remote)

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("sorcerer: Hostname error ", err)
	}
	fmt.Fprintf(w, "Hello %s from %s\n", remote, hostname)
}

//type Solution struct {
//Angle string
//Mass  string
//}

//func calculateSolution(w http.ResponseWriter, r *http.Request) {

//remote := r.RemoteAddr
//if r.Header.Get("X-FORWARDED-FOR") != "" {
//remote = r.Header.Get("X-FORWARDED-FOR")
//}

//log.Printf("sorcerer: handling request %s from %s", r.URL.Path, remote)

//solution := Solution{"23degrees", "10stone"}
//payload, err := json.Marshal(solution)
//if err != nil {
//log.Fatal("sorcerer: Payload error ", err)
//http.Error(w, err.Error(), http.StatusInternalServerError)
//return
//}

//w.Header().Set("Content-Type", "application/json")
//w.Write(payload)
//}

type sorcererService struct{}

func (s *sorcererService) FiringSolution(ctx context.Context, in *Target) (*Solution, error) {
	log.Printf("sorcerer: received request for a target distance %0.2f in %0.2f wind", in.Distance, in.Wind)
	return &Solution{Angle: 24.4, Mass: 10.0}, nil
}

func main() {
	sorcererPort := os.Getenv("SORCERER_PORT")
	if sorcererPort == "" {
		sorcererPort = "5000"
	}

	lis, err := net.Listen("tcp", ":"+sorcererPort)
	if err != nil {
		log.Fatal("sorcerer: Failed to listen on port... error: ", err)
	}

	server := grpc.NewServer()
	RegisterSorcererServiceServer(server, &sorcererService{})
	if err := server.Serve(lis); err != nil {
		log.Fatal(errors.Wrap(err, "sorcerer: Failed to start server!"))
	}
}
