package iyzipay

import (
	"encoding/json"
	"github.com/yahya077/iyzipay-go/client"
	"github.com/yahya077/iyzipay-go/requestSchema"
	"github.com/yahya077/iyzipay-go/responseSchema"
)

const (
	API_URL     = "https://api.iyzipay.com"
	SANDBOX_URL = "https://sandbox-api.iyzipay.com"
)

type Iyzipay struct {
	Client client.IyzicoClient
}

func (p Iyzipay) CreatePayment(request requestSchema.CreatePayment) (res responseSchema.CreatePayment, error error) {
	r, e := p.Client.Post(request, "/payment/auth")

	if e != nil {
		return res, e
	}

	e = json.NewDecoder(r.Body).Decode(&res)

	return res, e
}

func New(iyzicoClient client.IyzicoClient) Iyzipay {
	return Iyzipay{
		Client: iyzicoClient,
	}
}
