package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"jdramirezl/proyecto-1-topicos/mom/pkg/client"
)

func GetVar(r *http.Request, key string) (value string) {
	return r.URL.Query().Get(key)
}

func CheckError(
	err error, 
	printError string, 
	httpError string,
	w http.ResponseWriter, 
	){
	if err != nil {
		log.Printf(printError + "%v", err)
		http.Error(w, httpError, http.StatusBadRequest)
		return
	}
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
	mux.HandleFunc("/removeConnection", func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		CheckError(err, "Error reading body: %v", )
		if err != nil {
			log.Printf(, err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		connectionIP := r.URL.Query().Get("IP")
		err = momClient.AddConnection(connectionIP)

		if err != nil {
			log.Printf("Error connecting: %v", err)
			http.Error(w, "Can't connect", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

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
		if brokerType == "queue" {
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
