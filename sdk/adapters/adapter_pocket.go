package sdkAdapters

import (
	"fmt"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	pocket "github.com/techpartners-asia/pocket-go"
)

// PocketAdapter implements PaymentProvider for Pocket.
type PocketAdapter struct {
	client pocket.Pocket
}

func NewPocketAdapter(input types.PocketAdapter) *PocketAdapter {
	return &PocketAdapter{client: pocket.New(input.Merchant, input.ClientID, input.ClientSecret, input.Environment, input.TerminalIDRaw)}
}

func (a *PocketAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("pocket adapter not configured")
	}

	req := pocket.PocketCreateInvoiceInput{
		Amount:      input.Amount,
		OrderNumber: input.UID,
		InvoiceType: "ZERO",
		Channel:     "merchant",
		Info:        input.Note,
	}

	res, err := a.client.CreateInvoice(req)
	if err != nil {
		return nil, err
	}

	return &types.InvoiceResult{
		BankInvoiceID: fmt.Sprintf("%d", res.ID),
		BankQRCode:    res.Qr,
		Deeplinks: []types.Deeplink{{
			Name:        "Pocket",
			Description: "Pocket",
			Link:        res.DeepLink,
		}},
		IsPaid: false,
		Raw:    res,
	}, nil
}

func (a *PocketAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("pocket adapter not configured")
	}

	res, err := a.client.GetInvoiceByOrderNumber(input.UID)
	if err != nil {
		return nil, err
	}

	return &types.CheckInvoiceResult{
		IsPaid: res.State == "paid",
	}, nil
}
