package services

import (
	"errors"
	"fmt"
	"jamq-replica/constants"
	"jamq-replica/v1/exceptions"
)

/*
   The queue is implemented in the form of a doubly linked list
   This way we can easily consume from the front and
   insert into the back using 2 pointers.
*/

type Pair struct {
	First  interface{}
	Second interface{}
}

type FifoNode struct {
	Message   string
	Timestamp int
	Next      *FifoNode
	Prev      *FifoNode
}

type Queue interface {
	CreateQueue(string) error
	SendMessage(string, string, int) error
	ConsumeMessage(string) (*FifoNode, error)
	PurgeQueue(string) error
	DoesQueueExist(string) bool
}

var fifos map[string]*Pair = make(map[string]*Pair)

var fifosCounts map[string]int = make(map[string]int)

type fifoQueue struct {
}

type queueFactory struct {
}

var FifoQueue *fifoQueue = new(fifoQueue)
var QueueFactory *queueFactory = new(queueFactory)

func (qf *queueFactory) GetInstance(queueType string) (interface{}, error) {
	queueTypes := constants.GetQueueTypes()
	switch queueType {
	case queueTypes.FIFO:
		return FifoQueue, nil
	default:
		return nil, errors.New("Invalid queue type")
	}
}

func (fq *fifoQueue) DoesQueueExist(queueName string) bool {
	if _, exists := fifosCounts[queueName]; exists {
		return true
	}
	return false
}

func (fq *fifoQueue) CreateQueue(queueName string) error {
	if fq.DoesQueueExist(queueName) {
		return new(exceptions.QueueAlreadyExistsError)
	}
	fifosCounts[queueName] = 0
	return nil
}

func (fq *fifoQueue) SendMessage(queueName string, message string, timestamp int) error {
	if !fq.DoesQueueExist(queueName) {
		return new(exceptions.QueueNotFoundError)
	}
	/*
	   doubly linked list insertion if nothing was previously present or if queue was empty.
	*/
	if fifos[queueName] == nil || fifos[queueName].First == nil {
		newNode := &FifoNode{Message: message, Timestamp: timestamp, Next: nil, Prev: nil}
		if fifos[queueName] == nil {
			fifos[queueName] = &Pair{First: newNode, Second: newNode}
		} else {
			fifos[queueName].First = newNode
			fifos[queueName].Second = newNode
		}
		fifosCounts[queueName]++
		return nil
	}
	if fifos[queueName].First == nil {
		fmt.Println("First is null")
	}
	if fifos[queueName].Second == nil {
		fmt.Println("Second is null")
	}
	/*
	   doubly linked list insertion at the end
	*/
	pointers := fifos[queueName]
	newNode := &FifoNode{Message: message, Timestamp: timestamp, Next: nil, Prev: pointers.Second.(*FifoNode)}
	pointers.Second.(*FifoNode).Next = newNode
	pointers.Second = newNode
	fifosCounts[queueName]++
	return nil
}

func (fq *fifoQueue) ConsumeMessage(queueName string) (*FifoNode, error) {
	if fifos[queueName] == nil || fifos[queueName].First == nil {
		return nil, new(exceptions.QueueEmptyError)
	}
	currentNode := fifos[queueName].First.(*FifoNode)
	if currentNode.Next == nil {
		fifos[queueName].First = nil
		fifos[queueName].Second = nil
	} else {
		fifos[queueName].First = currentNode.Next
		currentNode.Next.Prev = nil
		fifosCounts[queueName]--
	}
	return currentNode, nil
}

func (fq *fifoQueue) PurgeQueue(queueName string) error {
	if fifos[queueName] == nil {
		return new(exceptions.QueueNotFoundError)
	}
	fifos[queueName].First = nil
	fifos[queueName].Second = nil
	return nil
}
