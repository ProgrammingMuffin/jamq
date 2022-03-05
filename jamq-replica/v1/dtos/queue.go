package dtos

type SendMessageResponse struct {
	Response string `json:"response"`
}

type ConsumeMessageResponse struct {
	Message   string `json:"message"`
	MessageId string `json:"message_id"`
}

type CreateQueueResponse struct {
	Response string `json:"response"`
}

type PurgeQueueResponse struct {
	Response string `json:"response"`
}
