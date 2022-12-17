package responseSchema

type Base struct {
	Status         string `json:"status"`
	ErrorCode      string `json:"errorCode"`
	ErrorMessage   string `json:"errorMessage"`
	ErrorGroup     string `json:"errorGroup"`
	Locale         string `json:"locale"`
	SystemTime     int    `json:"systemTime"`
	ConversationId string `json:"conversationId"`
}
