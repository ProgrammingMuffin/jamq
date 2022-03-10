package dtos

type CommonQueueRequestData struct {
	QueueName string `json:"queue_name"`
	QueueType string `json:"queue_type"`
}

type CreateQueueRequest struct {
	*CommonQueueRequestData
}

type PurgeQueueRequest struct {
	*CommonQueueRequestData
}

type SendMessageRequest struct {
	*CommonQueueRequestData
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}

type ReceiveMessageRequest struct {
	*CommonQueueRequestData
}
