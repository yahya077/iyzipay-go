package responseSchema

type CreatePayment struct {
	AuthCode                     string            `json:"authCode"`
	BasketID                     string            `json:"basketId"`
	BinNumber                    string            `json:"binNumber"`
	CardAssociation              string            `json:"cardAssociation"`
	CardFamily                   string            `json:"cardFamily"`
	CardType                     string            `json:"cardType"`
	ConversationID               string            `json:"conversationId"`
	Currency                     string            `json:"currency"`
	FraudStatus                  int64             `json:"fraudStatus"`
	HostReference                string            `json:"hostReference"`
	Installment                  int64             `json:"installment"`
	ItemTransactions             []ItemTransaction `json:"itemTransactions"`
	IyziCommissionFee            float64           `json:"iyziCommissionFee"`
	IyziCommissionRateAmount     float64           `json:"iyziCommissionRateAmount"`
	LastFourDigits               string            `json:"lastFourDigits"`
	Locale                       string            `json:"locale"`
	MdStatus                     int64             `json:"mdStatus"`
	MerchantCommissionRate       float64           `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64           `json:"merchantCommissionRateAmount"`
	PaidPrice                    float64           `json:"paidPrice"`
	PaymentID                    string            `json:"paymentId"`
	Phase                        string            `json:"phase"`
	Price                        float64           `json:"price"`
	Status                       string            `json:"status"`
	SystemTime                   int64             `json:"systemTime"`
}
