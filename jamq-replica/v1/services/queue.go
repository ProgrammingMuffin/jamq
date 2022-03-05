package services

import (
	"errors"
	"jamq-replica/constants"
	"jamq-replica/v1/exceptions"
)

type FifoNode struct {
	Message   string
	MessageId string
	Next      *FifoNode
}

type queue interface {
	createQueue(string)
	sendMessage(string, string)
	consumeMessage(string) string
	purgeQueue(string)
}

var fifos map[string]*FifoNode = make(map[string]*FifoNode)

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

func (fq *fifoQueue) createQueue(queueName string) error {
	if _, exists := fifosCounts[queueName]; exists {
		return new(exceptions.QueueAlreadyExistsError)
	}
	fifosCounts[queueName] = 0
	return nil
}

func (fq *fifoQueue) sendMessage(queueName string, message string) {
}

func (fq *fifoQueue) consumeMessage(queueName string, message string) {
}

func (fq *fifoQueue) purgeQueue(queueName string) {
}
