package requestSchema

type CreateCard struct {
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
	ExternalId     string `json:"externalId"`
	Email          string `json:"email"`
	CardUserKey    string `json:"cardUserKey"`
	Card           Card   `json:"card"`
}

type Card struct {
	CardAlias      string `json:"cardAlias"`
	CardNumber     string `json:"cardNumber"`
	ExpireYear     string `json:"expireYear"`
	ExpireMonth    string `json:"expireMonth"`
	CardHolderName string `json:"cardHolderName"`
}
