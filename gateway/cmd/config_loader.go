package main

import (
	"bufio"
	"os"
	"strings"
)

type Configuration struct {
	APIPort    string
	APIIP      string
	GRPCPort   string
	GRPCIP     string
	RabbitPort string
	RabbitIP   string
	RabbitQ    string
}

func loadConfig(directory string) (*Configuration, error) {
	file, err := os.Open(directory + "/env")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Configuration{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		switch parts[0] {
		case "API_PORT":
			config.APIPort = parts[1]
		case "API_IP":
			config.APIIP = parts[1]
		case "GRPC_PORT":
			config.GRPCPort = parts[1]
		case "GRPC_IP":
			config.GRPCIP = parts[1]
		case "RABBIT_PORT":
			config.RabbitPort = parts[1]
		case "RABBIT_IP":
			config.RabbitIP = parts[1]
		case "RABBIT_QUEUE":
			config.RabbitQ = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}