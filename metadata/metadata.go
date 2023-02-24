package metadata

type Format string

const (
    FormatBasic  Format = "basic"
    FormatPayout Format = "card_payouts"
)

type SourceType string

const (
    SourceTypeCard  = "card"
    SourceTypeToken = "token"
    SourceTypeBin   = "bin"
)

type Source struct {
    Type   string `json:"type"`
    Token  string `json:"token,omitempty"`
    Number string `json:"number,omitempty"`
    Bin    string `json:"bin,omitempty"`
}

type Request struct {
    Source Source `json:"source"`
    Format Format `json:"format"`
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
