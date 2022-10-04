package vmc

type Subscription struct {
	AnniversaryBillingDate string `json:"anniversary_billing_date"`
	AutoRenewedAllowed     string `json:"auto_renewed_allowed"`
	BillingFrequency       string `json:"billing_frequency"`
	BillingSubscriptionID  string `json:"billing_subscription_id"`
	CommitmentTerm         string `json:"commitment_term"`
	CommitmentTermUom      string `json:"commitment_term_uom"`
	CspSubscriptionID      string `json:"csp_subscription_id"`
	Description            string `json:"description"`
	EndDate                string `json:"end_date"`
	OfferName              string `json:"offer_name"`
	OfferType              string `json:"offer_type"`
	OfferVersion           string `json:"offer_version"`
	ProductID              string `json:"product_id"`
	ProductName            string `json:"product_name"`
	Quantity               string `json:"quantity"`
	Region                 string `json:"region"`
	StartDate              string `json:"start_date"`
	Status                 string `json:"status"`
}

type Subscriptions []Subscription

type OfferInstances struct {
	Offers []struct {
		CommitmentTerm int    `json:"commitment_term"`
		Currency       string `json:"currency"`
		Description    string `json:"description"`
		Name           string `json:"name"`
		Product        string `json:"product"`
		Region         string `json:"region"`
		Type           string `json:"type"`
		UnitPrice      string `json:"unit_price"`
		Version        string `json:"version"`
	} `json:"offers"`
	OnDemand struct {
		Currency    string `json:"currency"`
		Description string `json:"description"`
		MonthlyCost string `json:"monthly_cost"`
		Name        string `json:"name"`
		Product     string `json:"product"`
		Region      string `json:"region"`
		Type        string `json:"type"`
		UnitPrice   string `json:"unit_price"`
		Version     string `json:"version"`
	} `json:"on_demand"`
}

type SubscriptionProduct struct {
	Product string   `json:"product"`
	Types   []string `json:"types"`
}

type SubscriptionProducts []SubscriptionProduct
