package dtos

type SendMessageResponse struct {
	Response string `json:"response"`
}

type ConsumeMessageResponse struct {
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}

type CreateQueueResponse struct {
	Response string `json:"response"`
}

type PurgeQueueResponse struct {
	Response string `json:"response"`
}
