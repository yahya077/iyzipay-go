package responseSchema

type CreateCard struct {
	Base
	CardUserKey string `json:"cardUserKey"`
	CardDetail
}
