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

func NewSimpleAdapter(input types.SimpleAdapter) *SimpleAdapter {
	return &SimpleAdapter{client: simple.New(input.UserName, input.Password, input.BaseUrl, input.CallbackUrl)}
}

func (a *SimpleAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("simple adapter not configured")
	}

	expireAt := time.Now().Add(20 * time.Minute).Format("2006-01-02 15:04:05")

	req := simple.SimpleCreateInvoiceInput{
		OrderID:    input.UID,
		Total:      int(input.Amount),
		ExpireDate: expireAt,
	}

	res, err := a.client.CreateInvoice(req)
	if err != nil {
		return nil, err
	}

	return &types.InvoiceResult{
		BankInvoiceID: input.UID,
		Raw:           res,
		IsPaid:        false,
	}, nil
}

func (a *SimpleAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("simple adapter not configured")
	}

	res, err := a.client.GetInvoice(simple.SimpleGetInvoiceRequest{
		OrderID:  input.UID,
		SimpleID: "",
	})
	if err != nil {
		return nil, err
	}

	return &types.CheckInvoiceResult{
		IsPaid: res.Data.InvoiceStatus == "PAID",
	}, nil
}
