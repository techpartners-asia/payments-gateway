package types

import (
	"fmt"
	"strconv"
)

// PaymentType enumerates supported providers for Create().
type PaymentType string

const (
	PaymentTypeQPay     PaymentType = "qpay"
	PaymentTypeTokipay  PaymentType = "tokipay"
	PaymentTypeMonpay   PaymentType = "monpay"
	PaymentTypeGolomt   PaymentType = "golomt"
	PaymentTypeSocial   PaymentType = "socialpay"
	PaymentTypeStorePay PaymentType = "storepay"
	PaymentTypePocket   PaymentType = "pocket"
	PaymentTypeSimple   PaymentType = "simple"
	PaymentTypeBalc     PaymentType = "balc"
)

// InvoiceInput is the unified request accepted by CreateInvoice.
// Fill only the fields relevant to the chosen PaymentType.
type InvoiceInput struct {
	Type          PaymentType
	Amount        float64
	PaymentUID    string
	IsOrg         bool
	OrgRegNo      string
	Phone         string
	CustomerID    uint
	OrderUID      string
	CountryCode   string // optional, defaults to "+976" for Tokipay
	Note          string // optional description/notes
	CallbackURL   string // optional, used by some providers (e.g. Golomt)
	ReturnType    string // optional, used by some providers (e.g. Golomt) "GET" or "POST"
	ExpireMinutes int    // optional, used by Simple; defaults applied if zero
}

// Deeplink is a portable representation of provider deeplinks.
type Deeplink struct {
	Name        string
	Description string
	Link        string
	Logo        string
}

// InvoiceResult is a normalized result across providers.
type InvoiceResult struct {
	BankInvoiceID string
	BankQRCode    string
	Deeplinks     []Deeplink
	IsPaid        bool
	Raw           any
}

// PaymentProvider is the common interface adapters must implement.
type PaymentProvider interface {
	CreateInvoice(input InvoiceInput) (*InvoiceResult, error)
}

type (
	QPayConfig struct {
		Username    string
		Password    string
		Endpoint    string
		Callback    string
		InvoiceCode string
		MerchantID  string
	}
	TokipayConfig struct {
		Endpoint      string
		APIKey        string
		IMAPIKey      string
		Authorization string
		MerchantID    string
		SuccessURL    string
		FailureURL    string
		AppSchemaIOS  string
	}
	MonpayConfig struct {
		Endpoint  string
		Username  string
		AccountID string
		Callback  string
	}
	GolomtConfig struct {
		BaseURL     string
		Secret      string
		BearerToken string
	}
	SocialPayConfig struct {
		Terminal string
		Secret   string
		Endpoint string
	}
	StorePayConfig struct {
		AppUsername string
		AppPassword string
		Username    string
		Password    string
		AuthURL     string
		BaseURL     string
		StoreID     string
		CallbackURL string
	}
	PocketConfig struct {
		Merchant      string
		ClientID      string
		ClientSecret  string
		Environment   string
		TerminalIDRaw string // keep raw to avoid forcing caller to parse; parsed internally
	}
	SimpleConfig struct {
		Username    string
		Password    string
		BaseURL     string
		CallbackURL string
	}
	BalcAPIConfig struct {
		Endpoint string
		Token    string
	}
)

// Config groups provider configs so the SDK can be initialized in one call.
type Config struct {
	Qpay      QPayConfig
	TokiPay   TokipayConfig
	Monpay    MonpayConfig
	Golomt    GolomtConfig
	SocialPay SocialPayConfig
	StorePay  StorePayConfig
	Pocket    PocketConfig
	Simple    SimpleConfig
	BalcAPI   BalcAPIConfig
}

// QPayConfig helpers
func (c QPayConfig) Enabled() bool {
	return c.Username != "" || c.Password != "" || c.Endpoint != "" || c.Callback != "" || c.InvoiceCode != "" || c.MerchantID != ""
}
func (c QPayConfig) Validate() error {
	switch {
	case c.Username == "":
		return fmt.Errorf("qpay username is required")
	case c.Password == "":
		return fmt.Errorf("qpay password is required")
	case c.Endpoint == "":
		return fmt.Errorf("qpay endpoint is required")
	case c.Callback == "":
		return fmt.Errorf("qpay callback is required")
	case c.InvoiceCode == "":
		return fmt.Errorf("qpay invoice code is required")
	case c.MerchantID == "":
		return fmt.Errorf("qpay merchant id is required")
	default:
		return nil
	}
}

// TokipayConfig helpers
func (c TokipayConfig) Enabled() bool {
	return c.Endpoint != "" || c.APIKey != "" || c.IMAPIKey != "" || c.Authorization != "" || c.MerchantID != "" || c.SuccessURL != "" || c.FailureURL != "" || c.AppSchemaIOS != ""
}
func (c TokipayConfig) Validate() error {
	switch {
	case c.Endpoint == "":
		return fmt.Errorf("tokipay endpoint is required")
	case c.APIKey == "":
		return fmt.Errorf("tokipay api key is required")
	case c.IMAPIKey == "":
		return fmt.Errorf("tokipay im api key is required")
	case c.Authorization == "":
		return fmt.Errorf("tokipay authorization is required")
	case c.MerchantID == "":
		return fmt.Errorf("tokipay merchant id is required")
	case c.SuccessURL == "":
		return fmt.Errorf("tokipay success url is required")
	case c.FailureURL == "":
		return fmt.Errorf("tokipay failure url is required")
	case c.AppSchemaIOS == "":
		return fmt.Errorf("tokipay app schema ios is required")
	default:
		return nil
	}
}

// MonpayConfig helpers
func (c MonpayConfig) Enabled() bool {
	return c.Endpoint != "" || c.Username != "" || c.AccountID != "" || c.Callback != ""
}
func (c MonpayConfig) Validate() error {
	switch {
	case c.Endpoint == "":
		return fmt.Errorf("monpay endpoint is required")
	case c.Username == "":
		return fmt.Errorf("monpay username is required")
	case c.AccountID == "":
		return fmt.Errorf("monpay account id is required")
	case c.Callback == "":
		return fmt.Errorf("monpay callback is required")
	default:
		return nil
	}
}

// GolomtConfig helpers
func (c GolomtConfig) Enabled() bool {
	return c.BaseURL != "" || c.Secret != "" || c.BearerToken != ""
}
func (c GolomtConfig) Validate() error {
	switch {
	case c.BaseURL == "":
		return fmt.Errorf("golomt base url is required")
	case c.Secret == "":
		return fmt.Errorf("golomt secret is required")
	case c.BearerToken == "":
		return fmt.Errorf("golomt bearer token is required")
	default:
		return nil
	}
}

// SocialPayConfig helpers
func (c SocialPayConfig) Enabled() bool {
	return c.Terminal != "" || c.Secret != "" || c.Endpoint != ""
}
func (c SocialPayConfig) Validate() error {
	switch {
	case c.Terminal == "":
		return fmt.Errorf("socialpay terminal is required")
	case c.Secret == "":
		return fmt.Errorf("socialpay secret is required")
	case c.Endpoint == "":
		return fmt.Errorf("socialpay endpoint is required")
	default:
		return nil
	}
}

// StorePayConfig helpers
func (c StorePayConfig) Enabled() bool {
	return c.AppUsername != "" || c.AppPassword != "" || c.Username != "" || c.Password != "" || c.AuthURL != "" || c.BaseURL != "" || c.StoreID != "" || c.CallbackURL != ""
}
func (c StorePayConfig) Validate() error {
	switch {
	case c.AppUsername == "":
		return fmt.Errorf("storepay app username is required")
	case c.AppPassword == "":
		return fmt.Errorf("storepay app password is required")
	case c.Username == "":
		return fmt.Errorf("storepay username is required")
	case c.Password == "":
		return fmt.Errorf("storepay password is required")
	case c.AuthURL == "":
		return fmt.Errorf("storepay auth url is required")
	case c.BaseURL == "":
		return fmt.Errorf("storepay base url is required")
	case c.StoreID == "":
		return fmt.Errorf("storepay store id is required")
	case c.CallbackURL == "":
		return fmt.Errorf("storepay callback url is required")
	default:
		return nil
	}
}

// PocketConfig helpers
func (c PocketConfig) Enabled() bool {
	return c.Merchant != "" || c.ClientID != "" || c.ClientSecret != "" || c.Environment != "" || c.TerminalIDRaw != ""
}
func (c PocketConfig) Validate() (int64, error) {
	switch {
	case c.Merchant == "":
		return 0, fmt.Errorf("pocket merchant is required")
	case c.ClientID == "":
		return 0, fmt.Errorf("pocket client id is required")
	case c.ClientSecret == "":
		return 0, fmt.Errorf("pocket client secret is required")
	case c.Environment == "":
		return 0, fmt.Errorf("pocket environment is required")
	case c.TerminalIDRaw == "":
		return 0, fmt.Errorf("pocket terminal id is required")
	default:
		id, err := strconv.ParseInt(c.TerminalIDRaw, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("pocket terminal id parse: %w", err)
		}
		return id, nil
	}
}

// SimpleConfig helpers
func (c SimpleConfig) Enabled() bool {
	return c.Username != "" || c.Password != "" || c.BaseURL != "" || c.CallbackURL != ""
}
func (c SimpleConfig) Validate() error {
	switch {
	case c.Username == "":
		return fmt.Errorf("simple username is required")
	case c.Password == "":
		return fmt.Errorf("simple password is required")
	case c.BaseURL == "":
		return fmt.Errorf("simple base url is required")
	case c.CallbackURL == "":
		return fmt.Errorf("simple callback url is required")
	default:
		return nil
	}
}

// BalcAPIConfig helpers
func (c BalcAPIConfig) Enabled() bool {
	return c.Endpoint != "" || c.Token != ""
}
func (c BalcAPIConfig) Validate() error {
	switch {
	case c.Endpoint == "":
		return fmt.Errorf("balcapi endpoint is required")
	case c.Token == "":
		return fmt.Errorf("balcapi token is required")
	default:
		return nil
	}
}
