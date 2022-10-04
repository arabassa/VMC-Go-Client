package vmc

type Orgs []Org

type Org struct {
	UserId            string       `json:"user_id"`
	UserName          string       `json:"user_name"`
	Created           string       `json:"created"`
	Version           int          `json:"version"`
	Id                string       `json:"id"`
	UpdatedByUserId   string       `json:"updated_by_user_id"`
	UpdatedByUserName string       `json:"updated_by_user_name"`
	Updated           string       `json:"updated"`
	Name              string       `json:"name"`
	DisplayName       string       `json:"display_name"`
	ProjectState      string       `json:"project_state"`
	Properties        OrgProps     `json:"properties"`
	Type              string       `json:"org_type"`
	OrgSellerInfo     OrgSellerInf `json:"org_seller_info"`
}

type OrgSellerInf struct {
	Seller          string `json:"seller"`
	SellerAccountId string `json:"seller_account_id"`
}

type OrgProps struct {
	Values map[string]string `json:"values"`
}

type PaymentMethodInfo struct {
	DefaultFlag     bool   `json:"default_flag"`
	PaymentMethodID string `json:"payment_method_id"`
	Type            string `json:"type"`
}

type PaymentMethodInfos []PaymentMethodInfo

type AwsCloudProvider struct {
	Provider string   `json:"provider"`
	Regions  []string `json:"regions"`
}

type AwsCloudProviders []AwsCloudProvider
