package requestSchema

type CreatePayment struct {
	Locale          string       `json:"locale"`
	ConversationId  string       `json:"conversationId"`
	Price           string       `json:"price"`
	PaidPrice       string       `json:"paidPrice"`
	BasketId        string       `json:"basketId"`
	PaymentGroup    string       `json:"paymentGroup"`
	PaymentCard     PaymentCard  `json:"paymentCard"`
	Buyer           Buyer        `json:"buyer"`
	BillingAddress  Address      `json:"shippingAddress"`
	ShippingAddress Address      `json:"billingAddress"`
	BasketItems     []BasketItem `json:"basketItems"`
	Currency        string       `json:"currency"`
}
