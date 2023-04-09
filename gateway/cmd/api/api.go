package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	client "github.com/jdramirezl/proyecto-1-topicos/mom/pkg/client"
)

func GetVar(r *http.Request, key string) (value string) {
	return r.URL.Query().Get(key)
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	hostIP := os.Getenv("MOM_HOST")
	hostPort := os.Getenv("MOM_PORT")
	momClient := client.NewClient(hostIP, hostPort)

	mux.HandleFunc("/addMessagingSystem", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/removeMessagingSystem", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/addSubscriber", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/removeSubscriber", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/addConnection", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/removeConnection", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/addMessage", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		message := string(body)
		broker := r.URL.Query().Get("broker")
		brokerType := r.URL.Query().Get("type")
		if brokerType == "rabbitmq" {
			err := momClient.PublishQueue(broker, message)
			if err != nil {
				log.Printf("Error publishing message: %v", err)
				http.Error(w, "can't publish message", http.StatusBadRequest)
				return
			}
		} else {
			err := momClient.PublishTopic(broker, message)
			if err != nil {
				log.Printf("Error publishing message: %v", err)
				http.Error(w, "can't publish message", http.StatusBadRequest)
				return
			}
		}
		w.WriteHeader(http.StatusCreated)
	})

	return mux
}

func RunHttp(listenAddr string) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: NewRouter(), // Our own instance of servemux
	}

	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
}
