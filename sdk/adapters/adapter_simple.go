package sdkAdapters

import (
	"fmt"
	"time"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	simple "github.com/techpartners-asia/simple-go"
)

// SimpleAdapter implements PaymentProvider for Simple.
type SimpleAdapter struct {
	client simple.Simple
}

func NewSimpleAdapter(client simple.Simple) *SimpleAdapter {
	return &SimpleAdapter{client: client}
}

func (a *SimpleAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("simple adapter not configured")
	}

	expMinutes := input.ExpireMinutes
	if expMinutes <= 0 {
		expMinutes = 20
	}
	expireAt := time.Now().Add(time.Duration(expMinutes) * time.Minute).Format("2006-01-02 15:04:05")

	req := simple.SimpleCreateInvoiceInput{
		OrderID:    input.PaymentUID,
		Total:      int(input.Amount),
		ExpireDate: expireAt,
	}

	res, err := a.client.CreateInvoice(req)
	if err != nil {
		return nil, err
	}

	return &types.InvoiceResult{
		BankInvoiceID: input.PaymentUID,
		Raw:           res,
		IsPaid:        false,
	}, nil
}

func (a *SimpleAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("simple adapter not configured")
	}

	res, err := a.client.GetInvoice(simple.SimpleGetInvoiceRequest{
		OrderID:  input.PaymentUID,
		SimpleID: "",
	})
	if err != nil {
		return nil, err
	}

	return &types.CheckInvoiceResult{
		IsPaid: res.Data.InvoiceStatus == "PAID",
	}, nil
}
