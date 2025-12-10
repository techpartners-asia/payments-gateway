package paymentssdk

import (
	"fmt"

	sdkAdapters "github.com/techpartners-asia/payments-gateway/sdk/adapters"
	"github.com/techpartners-asia/payments-gateway/sdk/types"
)

// Gateway bundles adapters and exposes a unified CreateInvoice entrypoint.
type Gateway struct {
	QPayAdapter       types.PaymentProvider
	TokiPayAdapter    types.PaymentProvider
	BalcCreditAdapter types.PaymentProvider
	GolomtAdapter     types.PaymentProvider
	SocialPayAdapter  types.PaymentProvider
	StorePayAdapter   types.PaymentProvider
	PocketAdapter     types.PaymentProvider
	SimpleAdapter     types.PaymentProvider
	MonpayAdapter     types.PaymentProvider
}

// NewGatewayFromSDK builds a Gateway from an already initialized SDK clients.
// Any missing client leaves the corresponding adapter nil.
func NewGatewayFromSDK(s *SDK) *Gateway {
	g := &Gateway{}
	if s == nil {
		return g
	}
	if s.Qpay != nil {
		g.QPayAdapter = sdkAdapters.NewQPayAdapter(s.Qpay)
	}
	if s.TokiPay != nil {
		g.TokiPayAdapter = sdkAdapters.NewTokiPayAdapter(s.TokiPay)
	}
	if s.BalcAPI != nil {
		g.BalcCreditAdapter = sdkAdapters.NewBalcCreditAdapter(s.BalcAPI)
	}
	if s.Golomt != nil {
		g.GolomtAdapter = sdkAdapters.NewGolomtAdapter(s.Golomt)
	}
	if s.SocialPay != nil {
		g.SocialPayAdapter = sdkAdapters.NewSocialPayAdapter(s.SocialPay)
	}
	if s.StorePay != nil {
		g.StorePayAdapter = sdkAdapters.NewStorePayAdapter(s.StorePay)
	}
	if s.Pocket != nil {
		g.PocketAdapter = sdkAdapters.NewPocketAdapter(s.Pocket)
	}
	if s.Simple != nil {
		g.SimpleAdapter = sdkAdapters.NewSimpleAdapter(s.Simple)
	}
	if s.Monpay != nil {
		g.MonpayAdapter = sdkAdapters.NewMonpayAdapter(s.Monpay)
	}
	return g
}

// NewGatewayWithConfig builds a Gateway by first constructing SDK clients from config.
func New(cfg types.Config) (*Gateway, error) {
	s, err := NewWithConfig(cfg)
	if err != nil {
		return nil, err
	}
	return NewGatewayFromSDK(s), nil
}

// CreateInvoice routes to the correct provider adapter.
func (g *Gateway) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if g == nil {
		return nil, fmt.Errorf("gateway is nil")
	}
	var provider types.PaymentProvider
	switch input.Type {
	case types.PaymentTypeQPay:
		provider = g.QPayAdapter
	case types.PaymentTypeTokipay:
		provider = g.TokiPayAdapter
	case types.PaymentTypeBalc:
		provider = g.BalcCreditAdapter
	case types.PaymentTypeGolomt:
		provider = g.GolomtAdapter
	case types.PaymentTypeSocial:
		provider = g.SocialPayAdapter
	case types.PaymentTypeStorePay:
		provider = g.StorePayAdapter
	case types.PaymentTypePocket:
		provider = g.PocketAdapter
	case types.PaymentTypeSimple:
		provider = g.SimpleAdapter
	case types.PaymentTypeMonpay:
		provider = g.MonpayAdapter
	default:
		return nil, fmt.Errorf("unsupported payment type: %s", input.Type)
	}
	if provider == nil {
		return nil, fmt.Errorf("adapter for %s is not configured", input.Type)
	}
	return provider.CreateInvoice(input)
}

func (g *Gateway) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if g == nil {
		return nil, fmt.Errorf("gateway is nil")
	}
	var provider types.PaymentProvider
	switch input.Type {
	case types.PaymentTypeQPay:
		provider = g.QPayAdapter
	case types.PaymentTypeTokipay:
		provider = g.TokiPayAdapter
	case types.PaymentTypeBalc:
		provider = g.BalcCreditAdapter
	case types.PaymentTypeGolomt:
		provider = g.GolomtAdapter
	case types.PaymentTypeSocial:
		provider = g.SocialPayAdapter
	case types.PaymentTypeStorePay:
		provider = g.StorePayAdapter
	case types.PaymentTypePocket:
		provider = g.PocketAdapter
	case types.PaymentTypeSimple:
		provider = g.SimpleAdapter
	case types.PaymentTypeMonpay:
		provider = g.MonpayAdapter
	default:
		return nil, fmt.Errorf("unsupported payment type: %s", input.Type)
	}
	if provider == nil {
		return nil, fmt.Errorf("adapter for %s is not configured", input.Type)
	}
	return provider.CheckInvoice(input)
}
