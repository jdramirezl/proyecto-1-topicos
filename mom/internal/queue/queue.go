package queue

type Queue struct {
	Messages *LinkedList
	Consumers []Consumer
	CurrentConsumer int
	ConsumerMap map[string] chan string
}

func (q *Queue) roundRobin() {
	for !q.Consumers[q.CurrentConsumer].Available  {
		q.CurrentConsumer = (q.CurrentConsumer + 1) % len(q.Consumers)
	}
}

func (q *Queue) sendMessage() {
	message, err := q.Messages.Pop()
	if err != nil {
		return
	}
	q.Consumers[q.CurrentConsumer].ReceiveMessage(message)
	q.Consumers[q.CurrentConsumer].Available = false
}

func (q *Queue) unicast() {
	q.roundRobin()	
	q.sendMessage()
}

func (q* Queue) addConsumer() chan string {
	channel := make(chan string)
	cons := consumer.NewConsumer()
	q.ConsumerMap[cons.ID] = channel
}

func (q *Queue) RemoveConsumer(id string) {
	delete(q.ConsumerMap, id)
}


func (q *Queue) Send() {
	q.unicast()
}