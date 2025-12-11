package sdkAdapters

import (
	"fmt"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	"github.com/techpartners-asia/golomt-api-go/socialpay"
)

// SocialPayAdapter implements PaymentProvider for SocialPay.
type SocialPayAdapter struct {
	client socialpay.SocialPay
}

func NewSocialPayAdapter(input types.SocialPayAdapter) *SocialPayAdapter {
	return &SocialPayAdapter{client: socialpay.New(input.Terminal, input.Secret, input.Endpoint)}
}

func (a *SocialPayAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("socialpay adapter not configured")
	}

	res, err := a.client.CreateInvoiceQR(socialpay.InvoiceInput{
		Amount:  input.Amount,
		Invoice: input.UID,
	})
	if err != nil {
		return nil, err
	}

	// The library response does not include a QR text field; return raw and invoice id.
	return &types.InvoiceResult{
		BankInvoiceID: input.UID,
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *SocialPayAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("socialpay adapter not configured")
	}

	res, err := a.client.CheckInvoice(socialpay.InvoiceInput{
		Invoice: input.UID,
		Amount:  input.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &types.CheckInvoiceResult{
		IsPaid: res.ResponseCode == "00",
	}, nil
}
