package common;

import (
	
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

func connect(){
	// Create a connection to the MOM server
}

func disconnect(){
	// Disconect from MOM server
}