package main

import (
	"log"
	"net"
	"os"

	"jdramirezl/proyecto-1-topicos/gateway/cmd/api"
)

func main() {

	apiAddr := net.JoinHostPort("0.0.0.0", os.Getenv("PORT"))
	if err := api.RunHttp(apiAddr); err != nil {
		log.Fatal(err)
	}
}
