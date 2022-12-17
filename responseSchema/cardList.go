package responseSchema

type CardList struct {
	Base
	CardDetails []CardDetail `json:"cardDetails"`
}
