package requestSchema

type DeleteCard struct {
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
	CardUserKey    string `json:"cardUserKey"`
	CardToken      string `json:"cardToken"`
}
