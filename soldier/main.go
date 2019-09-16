package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//"sorcerer"
)

type sorcererService struct{}

func main() {
	sorcererPort := os.Getenv("SORCERER_PORT")
	if sorcererPort == "" {
		sorcererPort = "sorcerer"
	}

	sorcererHost := os.Getenv("SORCERER_HOST")
	if sorcererHost == "" {
		sorcererHost = "5000"
	}

	address := sorcererHost + ":" + sorcererPort

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("soldier: Failed to connect: ", err)
	}
	defer conn.Close()

	log.Printf("Soldier: connected to %s", address)

	c := NewSorcererServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.FiringSolution(ctx, &Target{Distance: 60.0 /* m */, Wind: -10.0 /* m/s */})
	if err != nil {
		log.Fatal("soldier: Failed request for firing solution: ", err)
	}

	log.Printf("Solution: angle: %0.2f degrees, mass: %0.2f kg", r.Angle, r.Mass)
}
