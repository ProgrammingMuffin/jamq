package dtos

type SendMessageRequest struct {
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}

