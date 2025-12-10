package sdkAdapters

import (
	"fmt"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	golomt "github.com/techpartners-asia/golomt-api-go/ecommerce"
)

// GolomtAdapter implements PaymentProvider for Golomt ecommerce.
type GolomtAdapter struct {
	client golomt.GolomtEcommerce
}

func NewGolomtAdapter(client golomt.GolomtEcommerce) *GolomtAdapter {
	return &GolomtAdapter{client: client}
}

func (a *GolomtAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("golomt adapter not configured")
	}

	returnType := golomt.GET
	if input.ReturnType != "" {
		switch input.ReturnType {
		case "GET", "get":
			returnType = golomt.GET
		case "POST", "post":
			returnType = golomt.POST
		case "MOBILE", "mobile":
			returnType = golomt.MOBILE
		default:
			return nil, fmt.Errorf("invalid golomt return type: %s", input.ReturnType)
		}
	}

	req := golomt.CreateInvoiceInput{
		ReturnType:    returnType,
		Amount:        input.Amount,
		TransactionID: input.PaymentUID,
		Callback:      input.CallbackURL,
	}

	res, err := a.client.CreateInvoice(req)
	if err != nil {
		return nil, err
	}

	return &types.InvoiceResult{
		BankInvoiceID: res.Invoice,
		IsPaid:        false,
		Raw:           res,
	}, nil
}
