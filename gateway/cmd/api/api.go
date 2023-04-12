package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jdramirezl/proyecto-1-topicos/mom/pkg/client"
)

func GetVar(r *http.Request, key string) (value string) {
	return r.URL.Query().Get(key)
}

func CheckError(
	err error,
	printError string,
	httpError string,
	w http.ResponseWriter,
) bool {
	if err != nil {
		log.Printf(printError+"%v", err)
		http.Error(w, httpError, http.StatusBadRequest)
		return false
	}

	return true
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	hostIP := os.Getenv("MOM_HOST")
	hostPort := os.Getenv("MOM_PORT")
	selfIp := os.Getenv("SELF_IP")
	momClient := client.NewClient(hostIP, hostPort, selfIp)

	mux.HandleFunc("/createSystem", func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		safe := CheckError(err, "Error reading body: ", "can't read body", w)
		if !safe {
			return
		}

		clientIP := r.URL.Query().Get("IP")
		broker := r.URL.Query().Get("broker")
		brokerType := r.URL.Query().Get("type")

		err = nil
		if brokerType == "queue" {
			err = momClient.CreateQueue(broker, clientIP)
		} else {
			err = momClient.CreateTopic(broker, clientIP)
		}

		safe = CheckError(err, "Error creating system: ", "can't create system", w)
		if !safe {
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("/deleteSystem", func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		safe := CheckError(err, "Error reading body: ", "can't read body", w)
		if !safe {
			return
		}

		clientIP := r.URL.Query().Get("IP")
		broker := r.URL.Query().Get("broker")
		brokerType := r.URL.Query().Get("type")

		err = nil
		if brokerType == "queue" {
			err = momClient.DeleteQueue(broker, clientIP)
		} else {
			err = momClient.DeleteTopic(broker, clientIP)
		}

		safe = CheckError(err, "Error deleting system: ", "can't delete system", w)
		if !safe {
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		safe := CheckError(err, "Error reading body: ", "can't read body", w)
		if !safe {
			return
		}

		clientIP := r.URL.Query().Get("IP")
		broker := r.URL.Query().Get("broker")
		brokerType := r.URL.Query().Get("type")

		err = nil
		if brokerType == "queue" {
			err = momClient.SubscribeQueue(broker, clientIP)
		} else {
			err = momClient.SubscribeTopic(broker, clientIP)
		}

		safe = CheckError(err, "Error subscribing: ", "can't subscribe", w)
		if !safe {
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		safe := CheckError(err, "Error reading body: ", "can't read body", w)
		if !safe {
			return
		}

		clientIP := r.URL.Query().Get("IP")
		broker := r.URL.Query().Get("broker")
		brokerType := r.URL.Query().Get("type")

		err = nil
		if brokerType == "queue" {
			err = momClient.UnSubscribeQueue(broker, clientIP)
		} else {
			err = momClient.UnSubscribeTopic(broker, clientIP)
		}

		safe = CheckError(err, "Error unsubscribing: ", "can't unsubscribe", w)
		if !safe {
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("/addConnection", func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		safe := CheckError(err, "Error reading body: ", "can't read body", w)
		if !safe {
			return
		}

		connectionIP := r.URL.Query().Get("IP")
		fmt.Println("ConnectedIP: " + connectionIP)
		err = momClient.AddConnection(connectionIP)

		safe = CheckError(err, "Error connecting: ", "Can't connect", w)
		if !safe {
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("/removeConnection", func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		safe := CheckError(err, "Error reading body: ", "can't read body", w)
		if !safe {
			return
		}

		connectionIP := r.URL.Query().Get("IP")
		err = momClient.RemoveConnection(connectionIP)

		safe = CheckError(err, "Error connecting: ", "Can't connect", w)
		if !safe {
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

		safe := CheckError(err, "Error reading body: ", "can't read body", w)
		if !safe {
			return
		}

		message := string(body)
		broker := r.URL.Query().Get("broker")
		brokerType := r.URL.Query().Get("type")

		err = nil
		if brokerType == "queue" {
			err = momClient.PublishQueue(broker, message)
		} else {
			err = momClient.PublishTopic(broker, message)
		}

		safe = CheckError(err, "Error publishing message: ", "can't publish message", w)
		if !safe {
			return
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
