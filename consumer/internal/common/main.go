package common;

import (
	"encoding/json"
)

type Type int

const (
	QUEUE Type = iota
	TOPIC
)

type MessagingSystem struct {
	name string
	kind Type
	creator string
}

type Message struct {
	name string
	kind Type
	payload string
}

type Subscription struct {
	name string
	kind Type
	ip string
}

// function to marshal any struct to JSON bytes
func marshalJSON(v interface{}) ([]byte, error) {
    return json.Marshal(v)
}

// function to unmarshal JSON bytes to any struct
func unmarshalJSON(jsonBytes []byte, v interface{}) error {
    return json.Unmarshal(jsonBytes, v)
}

func connect(){
	// Create a connection to the MOM server
}

func disconnect(){
	// Disconect from MOM server
}