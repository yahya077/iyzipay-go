package api

import (
	"encoding/json"
	"github.com/yahya077/iyzipay-go/client"
	"github.com/yahya077/iyzipay-go/requestSchema"
	"github.com/yahya077/iyzipay-go/responseSchema"
)

type CardStorage struct {
	client client.IyzicoClient
}

func (p CardStorage) Get(request requestSchema.RetrieveCard) (res responseSchema.CardList, error error) {
	r, e := p.client.Post(request, "/cardstorage/cards")

	if e != nil {
		return res, e
	}

	e = json.NewDecoder(r.Body).Decode(&res)

	return res, e
}

func (p CardStorage) Create(request requestSchema.CreateCard) (res responseSchema.CreateCard, error error) {
	r, e := p.client.Post(request, "/cardstorage/card")

	if e != nil {
		return res, e
	}

	e = json.NewDecoder(r.Body).Decode(&res)

	return res, e
}

func (p CardStorage) Destroy(request requestSchema.DeleteCard) (res responseSchema.Base, error error) {
	r, e := p.client.Delete(request, "/cardstorage/card")

	if e != nil {
		return res, e
	}

	e = json.NewDecoder(r.Body).Decode(&res)

	return res, e
}

func NewCardStorage(client client.IyzicoClient) CardStorage {
	return CardStorage{client}
}
