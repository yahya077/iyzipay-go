package api

import (
	"encoding/json"
	"github.com/yahya077/iyzipay-go/client"
	"github.com/yahya077/iyzipay-go/requestSchema"
	"github.com/yahya077/iyzipay-go/responseSchema"
)

type Payment struct {
	client client.IyzicoClient
}

func (p Payment) Create(request requestSchema.CreatePayment) (res responseSchema.CreatePayment, error error) {
	r, e := p.client.Post(request, "/payment/auth")

	if e != nil {
		return res, e
	}

	e = json.NewDecoder(r.Body).Decode(&res)

	return res, e
}

func NewPayment(client client.IyzicoClient) Payment {
	return Payment{client}
}
