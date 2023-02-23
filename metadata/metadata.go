package metadata

type Request struct {
	Source struct {
		Type  string `json:"type"`
		Token string `json:"token"`
	} `json:"source"`
	Format string `json:"format"`
}

type Response struct {
	Bin               string `json:"bin"`
	Scheme            string `json:"scheme"`
	SchemeLocal       string `json:"scheme_local"`
	CardType          string `json:"card_type"`
	CardCategory      string `json:"card_category"`
	Issuer            string `json:"issuer"`
	IssuerCountry     string `json:"issuer_country"`
	IssuerCountryName string `json:"issuer_country_name"`
	ProductID         string `json:"product_id"`
	ProductType       string `json:"product_type"`
}
