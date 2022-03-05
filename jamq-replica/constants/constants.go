package constants

type headers struct {
	X_AUTH_SIGNED    string
	X_AUTH_SIGNATURE string
	X_AUTH_NONCE     string
}

type config struct {
	SECRET    string
	NONCE_TTL string
}

type queueTypes struct {
	FIFO string
}

var headersInstance *headers = nil
var configInstance *config = nil
var queueTypesInstance *queueTypes = nil

func GetHeaders() *headers {
	if headersInstance != nil {
		return headersInstance
	}
	headersInstance = &headers{}
	headersInstance.X_AUTH_NONCE = "X-AUTH-NONCE"
	headersInstance.X_AUTH_SIGNATURE = "X-AUTH-SIGNATURE"
	headersInstance.X_AUTH_SIGNED = "X-AUTH-SIGNED"
	return headersInstance
}

func GetConfig() *config {
	if configInstance != nil {
		return configInstance
	}
	configInstance.NONCE_TTL = "NONCE_TTL"
	configInstance.SECRET = "SECRET"
	return configInstance
}

func GetQueueTypes() *queueTypes {
	if queueTypesInstance != nil {
		return queueTypesInstance
	}
	queueTypesInstance := &queueTypes{}
	queueTypesInstance.FIFO = "FIFO"
	return queueTypesInstance
}
