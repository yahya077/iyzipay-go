package iyzipay

import (
	"github.com/yahya077/iyzipay-go/api"
	"github.com/yahya077/iyzipay-go/client"
)

const (
	API_URL     = "https://api.iyzipay.com"
	SANDBOX_URL = "https://sandbox-api.iyzipay.com"
)

type IIyzipay interface {
	CardStorage() api.CardStorage
	Payment() api.Payment
}

type Iyzipay struct {
	client client.IyzicoClient
}

func (p Iyzipay) CardStorage() api.CardStorage {
	return api.NewCardStorage(p.client)
}

func (p Iyzipay) Payment() api.Payment {
	return api.NewPayment(p.client)
}

func New(iyzicoClient client.IyzicoClient) Iyzipay {
	return Iyzipay{iyzicoClient}
}
