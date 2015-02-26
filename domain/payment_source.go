package domain

type PaymentSource struct {
	Id                       int64  `json:"id"`
	CCType                   string `json:"cc_type"`
	GatewayCustomerProfileId string `json:"gateway_customer_profile_id"`
	GatewayPaymentProfileId  string `json:"gateway_payment_profile_id"`
	Month                    string `json:"month"`
	Name                     string `json:"name"`
	LastDigits               string `json:"last_digits"`
	Year                     string `json:"year"`
}
