package sdkAdapters

import (
	"fmt"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	storepay "github.com/techpartners-asia/storepay-go"
)

// StorePayAdapter implements PaymentProvider for StorePay.
type StorePayAdapter struct {
	client storepay.Storepay
}

func NewStorePayAdapter(client storepay.Storepay) *StorePayAdapter {
	return &StorePayAdapter{client: client}
}

func (a *StorePayAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("storepay adapter not configured")
	}

	res, err := a.client.Loan(storepay.StorepayLoanInput{
		Amount:       input.Amount,
		MobileNumber: input.Phone,
		Description:  input.Note,
	})
	if err != nil {
		return nil, err
	}

	return &types.InvoiceResult{
		BankInvoiceID: fmt.Sprintf("%d", res),
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *StorePayAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("storepay adapter not configured")
	}

	res, err := a.client.LoanCheck(input.PaymentUID)
	if err != nil {
		return nil, err
	}

	return &types.CheckInvoiceResult{
		IsPaid: res,
	}, nil
}
