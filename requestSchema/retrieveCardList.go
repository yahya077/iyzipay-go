package requestSchema

type RetrieveCard struct {
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
	CardUserKey    string `json:"cardUserKey"`
}
