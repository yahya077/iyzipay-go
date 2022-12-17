package responseSchema

type CardDetail struct {
	BinNumber       string `json:"binNumber"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily      string `json:"cardFamily"`
	CardBankName    string `json:"cardBankName"`
	CardBankCode    string `json:"cardBankCode"`
	CardToken       string `json:"cardToken"`
	CardAlias       string `json:"cardAlias"`
}
