package paymentssdk

import (
	"fmt"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	balcapi "github.com/techpartners-asia/balc-api-go"
	golomt "github.com/techpartners-asia/golomt-api-go/ecommerce"
	"github.com/techpartners-asia/golomt-api-go/socialpay"
	"github.com/techpartners-asia/monpay-go/monpay"
	pocket "github.com/techpartners-asia/pocket-go"
	"github.com/techpartners-asia/qpay-go/qpay_v2"
	simple "github.com/techpartners-asia/simple-go"
	storepay "github.com/techpartners-asia/storepay-go"
	tokipay "github.com/techpartners-asia/tokipay-go"
)

// SDK bundles all supported payment providers behind a single initialization step.
type SDK struct {
	Qpay      qpay_v2.QPay
	TokiPay   tokipay.Tokipay
	Monpay    monpay.Monpay
	Golomt    golomt.GolomtEcommerce
	SocialPay socialpay.SocialPay
	StorePay  storepay.Storepay
	Pocket    pocket.Pocket
	Simple    simple.Simple
	BalcAPI   balcapi.Balc
}

// NewWithConfig initializes providers using the given Config. Any provider with
// an empty config is skipped; partial configs return an error so misconfigurations
// surface early.
func NewWithConfig(cfg types.Config) (*SDK, error) {
	var s SDK

	if cfg.Qpay.Enabled() {
		if err := cfg.Qpay.Validate(); err != nil {
			return nil, err
		}
		s.Qpay = qpay_v2.New(cfg.Qpay.Username, cfg.Qpay.Password, cfg.Qpay.Endpoint, cfg.Qpay.Callback, cfg.Qpay.InvoiceCode, cfg.Qpay.MerchantID)
	}
	if cfg.TokiPay.Enabled() {
		if err := cfg.TokiPay.Validate(); err != nil {
			return nil, err
		}
		s.TokiPay = tokipay.New(cfg.TokiPay.Endpoint, cfg.TokiPay.APIKey, cfg.TokiPay.IMAPIKey, cfg.TokiPay.Authorization, cfg.TokiPay.MerchantID, cfg.TokiPay.SuccessURL, cfg.TokiPay.FailureURL, cfg.TokiPay.AppSchemaIOS)
	}
	if cfg.Monpay.Enabled() {
		if err := cfg.Monpay.Validate(); err != nil {
			return nil, err
		}
		s.Monpay = monpay.New(cfg.Monpay.Endpoint, cfg.Monpay.Username, cfg.Monpay.AccountID, cfg.Monpay.Callback)
	}
	if cfg.Golomt.Enabled() {
		if err := cfg.Golomt.Validate(); err != nil {
			return nil, err
		}
		s.Golomt = golomt.New(cfg.Golomt.BaseURL, cfg.Golomt.Secret, cfg.Golomt.BearerToken)
	}
	if cfg.SocialPay.Enabled() {
		if err := cfg.SocialPay.Validate(); err != nil {
			return nil, err
		}
		s.SocialPay = socialpay.New(cfg.SocialPay.Terminal, cfg.SocialPay.Secret, cfg.SocialPay.Endpoint)
	}
	if cfg.StorePay.Enabled() {
		if err := cfg.StorePay.Validate(); err != nil {
			return nil, err
		}
		s.StorePay = storepay.New(cfg.StorePay.AppUsername, cfg.StorePay.AppPassword, cfg.StorePay.Username, cfg.StorePay.Password, cfg.StorePay.AuthURL, cfg.StorePay.BaseURL, cfg.StorePay.StoreID, cfg.StorePay.CallbackURL)
	}
	if cfg.Pocket.Enabled() {
		terminalID, err := cfg.Pocket.Validate()
		if err != nil {
			return nil, err
		}
		s.Pocket = pocket.New(cfg.Pocket.Merchant, cfg.Pocket.ClientID, cfg.Pocket.ClientSecret, cfg.Pocket.Environment, terminalID)
	}
	if cfg.Simple.Enabled() {
		if err := cfg.Simple.Validate(); err != nil {
			return nil, err
		}
		s.Simple = simple.New(cfg.Simple.Username, cfg.Simple.Password, cfg.Simple.BaseURL, cfg.Simple.CallbackURL)
	}
	if cfg.BalcAPI.Enabled() {
		if err := cfg.BalcAPI.Validate(); err != nil {
			return nil, err
		}
		s.BalcAPI = balcapi.New(cfg.BalcAPI.Endpoint, cfg.BalcAPI.Token)
	}

	return &s, nil
}

// Factories contains a constructor for each provider. Provide the ones you need.
// Each constructor should fully configure its client (credentials, endpoints, timeouts, etc.) before returning.
type Factories struct {
	Qpay      func() (qpay_v2.QPay, error)
	TokiPay   func() (tokipay.Tokipay, error)
	Monpay    func() (monpay.Monpay, error)
	Golomt    func() (golomt.GolomtEcommerce, error)
	SocialPay func() (socialpay.SocialPay, error)
	StorePay  func() (storepay.Storepay, error)
	Pocket    func() (pocket.Pocket, error)
	Simple    func() (simple.Simple, error)
	BalcAPI   func() (balcapi.Balc, error)
}

// NewFromFactories builds an SDK instance by invoking the provided factories.
func NewFromFactories(f Factories) (*SDK, error) {
	var (
		s   SDK
		err error
	)

	if f.Qpay != nil {
		if s.Qpay, err = f.Qpay(); err != nil {
			return nil, fmt.Errorf("init qpay: %w", err)
		}
	}
	if f.TokiPay != nil {
		if s.TokiPay, err = f.TokiPay(); err != nil {
			return nil, fmt.Errorf("init tokipay: %w", err)
		}
	}
	if f.Monpay != nil {
		if s.Monpay, err = f.Monpay(); err != nil {
			return nil, fmt.Errorf("init monpay: %w", err)
		}
	}
	if f.Golomt != nil {
		if s.Golomt, err = f.Golomt(); err != nil {
			return nil, fmt.Errorf("init golomt: %w", err)
		}
	}
	if f.SocialPay != nil {
		if s.SocialPay, err = f.SocialPay(); err != nil {
			return nil, fmt.Errorf("init socialpay: %w", err)
		}
	}
	if f.StorePay != nil {
		if s.StorePay, err = f.StorePay(); err != nil {
			return nil, fmt.Errorf("init storepay: %w", err)
		}
	}
	if f.Pocket != nil {
		if s.Pocket, err = f.Pocket(); err != nil {
			return nil, fmt.Errorf("init pocket: %w", err)
		}
	}
	if f.Simple != nil {
		if s.Simple, err = f.Simple(); err != nil {
			return nil, fmt.Errorf("init simple: %w", err)
		}
	}
	if f.BalcAPI != nil {
		if s.BalcAPI, err = f.BalcAPI(); err != nil {
			return nil, fmt.Errorf("init balcapi: %w", err)
		}
	}

	return &s, nil
}
