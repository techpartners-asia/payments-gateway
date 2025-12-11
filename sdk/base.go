package sdk

import (
	"fmt"

	sdkAdapters "github.com/techpartners-asia/payments-gateway/sdk/adapters"
	"github.com/techpartners-asia/payments-gateway/sdk/types"
)

type Input struct {
	Qpay      types.QpayAdapter
	TokiPay   types.TokipayAdapter
	StorePay  types.StorePayAdapter
	SocialPay types.SocialPayAdapter
	Simple    types.SimpleAdapter
	Pocket    types.PocketAdapter
	MonPay    types.MonpayAdapter
	Golomt    types.GolomtAdapter
	Balc      types.BalcAdapter
}

type SDK interface {
	Create(input types.InvoiceInput) (*types.InvoiceResult, error)
	Check(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error)
}

type sdk struct {
	input Input

	*sdkAdapters.BalcCreditAdapter
	*sdkAdapters.QPayAdapter
	*sdkAdapters.TokiPayAdapter
	*sdkAdapters.StorePayAdapter
	*sdkAdapters.SocialPayAdapter
	*sdkAdapters.SimpleAdapter
	*sdkAdapters.PocketAdapter
	*sdkAdapters.MonpayAdapter
	*sdkAdapters.GolomtAdapter
}

func New(input Input) SDK {
	return &sdk{
		input:             input,
		QPayAdapter:       sdkAdapters.NewQPayAdapter(input.Qpay),
		TokiPayAdapter:    sdkAdapters.NewTokiPayAdapter(input.TokiPay),
		StorePayAdapter:   sdkAdapters.NewStorePayAdapter(input.StorePay),
		SocialPayAdapter:  sdkAdapters.NewSocialPayAdapter(input.SocialPay),
		SimpleAdapter:     sdkAdapters.NewSimpleAdapter(input.Simple),
		PocketAdapter:     sdkAdapters.NewPocketAdapter(input.Pocket),
		MonpayAdapter:     sdkAdapters.NewMonpayAdapter(input.MonPay),
		GolomtAdapter:     sdkAdapters.NewGolomtAdapter(input.Golomt),
		BalcCreditAdapter: sdkAdapters.NewBalcCreditAdapter(input.Balc),
	}
}

func (s *sdk) Create(input types.InvoiceInput) (*types.InvoiceResult, error) {

	switch input.Type {
	case types.PaymentTypeQPay:
		return s.QPayAdapter.CreateInvoice(input)
	case types.PaymentTypeTokipay:
		return s.TokiPayAdapter.CreateInvoice(input)
	case types.PaymentTypeStorePay:
		return s.StorePayAdapter.CreateInvoice(input)
	case types.PaymentTypeSocial:
		return s.SocialPayAdapter.CreateInvoice(input)
	case types.PaymentTypeSimple:
		return s.SimpleAdapter.CreateInvoice(input)
	case types.PaymentTypePocket:
		return s.PocketAdapter.CreateInvoice(input)
	case types.PaymentTypeMonpay:
		return s.MonpayAdapter.CreateInvoice(input)
	case types.PaymentTypeGolomt:
		return s.GolomtAdapter.CreateInvoice(input)
	case types.PaymentTypeBalc:
		return s.BalcCreditAdapter.CreateInvoice(input)
	default:
		return nil, fmt.Errorf("unsupported payment type: %s", input.Type)
	}
}

func (s *sdk) Check(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	switch input.Type {
	case types.PaymentTypeQPay:
		return s.QPayAdapter.CheckInvoice(input)
	case types.PaymentTypeTokipay:
		return s.TokiPayAdapter.CheckInvoice(input)
	case types.PaymentTypeStorePay:
		return s.StorePayAdapter.CheckInvoice(input)
	case types.PaymentTypeSocial:
		return s.SocialPayAdapter.CheckInvoice(input)
	case types.PaymentTypeSimple:
		return s.SimpleAdapter.CheckInvoice(input)
	case types.PaymentTypePocket:
		return s.PocketAdapter.CheckInvoice(input)
	case types.PaymentTypeMonpay:
		return s.MonpayAdapter.CheckInvoice(input)
	case types.PaymentTypeGolomt:
		return s.GolomtAdapter.CheckInvoice(input)
	case types.PaymentTypeBalc:
		return s.BalcCreditAdapter.CheckInvoice(input)
	default:
		return nil, fmt.Errorf("unsupported payment type: %s", input.Type)
	}
}
