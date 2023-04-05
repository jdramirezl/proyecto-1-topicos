package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func bootstrap() *Configuration {
	fmt.Println("Loading configuration...")
	config, err := loadConfig("./config")
	failOnError(err, "Failed to load configuration")
	fmt.Println("Configuration loaded successfully")
	return config
}

func main() {
	// Start
	fmt.Println("Starting the gateway...")
	config := bootstrap()

	// Create addresses
	// TODO: CHANGE TO THE MASTER IP
	apiAddr := ":" + config.APIPort //net.JoinHostPort(config.APIIP, config.APIPort)
	grpcAddr := net.JoinHostPort(config.GRPCIP, config.GRPCPort)
	

	// Run GRPC
	connGRPC := runGRPC(grpcAddr)
	defer connGRPC.Close()

	// Run HTTP
	if err := RunHttp(apiAddr, connGRPC); err != nil {
		log.Fatal(err)
	}

}