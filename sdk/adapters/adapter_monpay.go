package sdkAdapters

import (
	"fmt"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	"github.com/techpartners-asia/monpay-go/monpay"
)

// MonpayAdapter currently does not implement invoice creation because the library
// exposes QR helpers rather than a direct create invoice API.
type MonpayAdapter struct {
	client monpay.Monpay
}

func NewMonpayAdapter(client monpay.Monpay) *MonpayAdapter {
	return &MonpayAdapter{client: client}
}

func (a *MonpayAdapter) CreateInvoice(types.InvoiceInput) (*types.InvoiceResult, error) {
	return nil, fmt.Errorf("monpay create invoice is not implemented; use monpay.GenerateQr or other helpers directly")
}

func (a *MonpayAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("monpay adapter not configured")
	}

	res, err := a.client.CheckQr(input.PaymentUID)
	if err != nil {
		return nil, err
	}

	return &types.CheckInvoiceResult{
		IsPaid: res.Code == 0,
	}, nil
}
