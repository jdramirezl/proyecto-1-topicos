package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc"
)

var (
	robin bool = true
)

func output(w http.ResponseWriter, response string, fail string, err error, len int) {
	// Error handling
	if err != nil {
		log.Fatalf("Error when calling func: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		log.Printf("Response from server list: %s", response)
		w.WriteHeader(http.StatusOK)

		if len == 0 {
			w.Write([]byte(fail))
		} else {
			w.Write([]byte(response))
		}
	}
}

func GetVar(r *http.Request, key string) (value string) {
	return r.URL.Query().Get(key)
}

func NewRouter(conn *grpc.ClientConn, chanRabbit *amqp.Channel, qRabbit amqp.Queue) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		var (
			files    []string
			response string
			err      error
		)

		// GRPC
		var resGRPC *FILERESPONSE
		log.Printf("Response from GRPC: list")
		resGRPC, err = List(conn)

		// Error Check
		failOnError(err, "Failed to call List")

		// Obtain result?
		files = resGRPC.Name
		
		// Join the files
		response = strings.Join(files, ", ")

		// Output
		output(w, response, "No files found", err, len(response))
	})

	mux.HandleFunc("/addMessagingSystem", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/removeMessagingSystem", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/addSubscriber", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/removeSubscriber", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/addConnection", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/removeConnection", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/addMessage", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/removeMessage", func(w http.ResponseWriter, r *http.Request) {})
	
	return mux
}

func RunHttp(listenAddr string, connGRPC *grpc.ClientConn) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: NewRouter(connGRPC), // Our own instance of servemux
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
}