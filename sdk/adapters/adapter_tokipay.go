package sdkAdapters

import (
	"fmt"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	tokipay "github.com/techpartners-asia/tokipay-go"
)

// TokiPayAdapter implements PaymentProvider for Tokipay.
type TokiPayAdapter struct {
	client tokipay.Tokipay
}

func NewTokiPayAdapter(client tokipay.Tokipay) *TokiPayAdapter {
	return &TokiPayAdapter{client: client}
}

func (a *TokiPayAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("tokipay adapter not configured")
	}
	country := input.CountryCode
	if country == "" {
		country = "+976"
	}
	res, err := a.client.PaymentSentUser(tokipay.TokipayPaymentInput{
		OrderId:     input.PaymentUID,
		Amount:      int64(input.Amount),
		PhoneNo:     input.Phone,
		CountryCode: country,
		Notes:       input.Note,
	})
	if err != nil {
		return nil, err
	}
	return &types.InvoiceResult{
		BankInvoiceID: input.PaymentUID,
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *TokiPayAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("tokipay adapter not configured")
	}

	res, err := a.client.PaymentStatus(input.PaymentUID)
	if err != nil {
		return nil, err
	}

	return &types.CheckInvoiceResult{
		IsPaid: res.Data.Status == "COMPLETED",
	}, nil
}
