package dtos

type CommonQueueRequestData struct {
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
