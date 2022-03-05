package exceptions

type QueueNotFoundError struct {
}

func (e *QueueNotFoundError) Error() string {
	return "The requested queue does not exist!"
}

type QueueAlreadyExistsError struct {
}

func (e *QueueAlreadyExistsError) Error() string {
	return "The requested queue already exists!"
}
