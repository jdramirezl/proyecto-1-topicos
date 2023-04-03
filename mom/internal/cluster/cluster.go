package cluster

import (

)


func join(){
	// Nodo nuevo envia IP a Nodo maestro
	// tambien a Coordinator
}

func heartbeat(){
	// Nodo en control envia heartbeat a los esclavos
	// Si esclavos no tienen heartbeat del control, hacen consenso
}

func caregiver(){
	// Check TTL of slaves, remove ones with TTL fromo config
	// TODO: Read TTL
}

func check(){
	// Check TTL of slaves, remove ones with TTL fromo config
	// TODO: Read TTL
}

func complain(){
	// When heartbeat is not received slaves create new master
}

func sendMessage(){
	// Master sends a message to all replicas
	// goroutines?
}

func removeMessage(){
	// Master says to remove already sent message
	// goroutines?
}

func sendChannel(){
	// Master sends a Channel to all replicas
	// goroutines?
}

func removeChannel(){
	// Master says to remove already sent Channel
	// goroutines?
}

func sendTopic(){
	// Master sends a Topic to all replicas
	// goroutines?
}

func removeTopic(){
	// Master says to remove already sent Topic
	// goroutines?
}

func sendSubscriber(){
	// Master sends a Topic to all replicas
	// goroutines?
}

func removeSubscriber(){
	// Master says to remove already sent Topic
	// goroutines?
}

func catchmeup(){
	// Llamado en JOIN
	// SOlicitar la info
}

func catchyouup(){
	// Llamado desde el main si nos solicitan info
	// Goroutine????
}