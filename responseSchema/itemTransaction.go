package responseSchema

type ItemTransaction struct {
	BlockageRate                  float64         `json:"blockageRate"`
	BlockageRateAmountMerchant    float64         `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float64         `json:"blockageRateAmountSubMerchant"`
	BlockageResolvedDate          string          `json:"blockageResolvedDate"`
	ConvertedPayout               ConvertedPayout `json:"convertedPayout"`
	ItemID                        string          `json:"itemId"`
	IyziCommissionFee             float64         `json:"iyziCommissionFee"`
	IyziCommissionRateAmount      float64         `json:"iyziCommissionRateAmount"`
	MerchantCommissionRate        float64         `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float64         `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          float64         `json:"merchantPayoutAmount"`
	PaidPrice                     float64         `json:"paidPrice"`
	PaymentTransactionID          string          `json:"paymentTransactionId"`
	Price                         float64         `json:"price"`
	SubMerchantPayoutAmount       float64         `json:"subMerchantPayoutAmount"`
	SubMerchantPayoutRate         float64         `json:"subMerchantPayoutRate"`
	SubMerchantPrice              float64         `json:"subMerchantPrice"`
	TransactionStatus             int64           `json:"transactionStatus"`
}

type ConvertedPayout struct {
	BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
	Currency                      string  `json:"currency"`
	IyziCommissionFee             float64 `json:"iyziCommissionFee"`
	IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
	IyziConversionRate            float64 `json:"iyziConversionRate"`
	IyziConversionRateAmount      float64 `json:"iyziConversionRateAmount"`
	MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
	PaidPrice                     float64 `json:"paidPrice"`
	SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
}
