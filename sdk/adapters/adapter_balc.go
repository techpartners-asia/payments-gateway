package sdkAdapters

import (
	"fmt"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	balcapi "github.com/techpartners-asia/balc-api-go"
)

// BalcCreditAdapter implements PaymentProvider for Balc credit flow.
type BalcCreditAdapter struct {
	client balcapi.Balc
}

func NewBalcCreditAdapter(input types.BalcAdapter) *BalcCreditAdapter {
	return &BalcCreditAdapter{client: balcapi.New(input.Endpoint, input.Token)}
}

func (a *BalcCreditAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("balc adapter not configured")
	}

	creditCheck, err := a.client.LimitCheck(int(input.CustomerID))
	if err != nil {
		return nil, fmt.Errorf("error on balcAPI check: %w", err)
	}
	if creditCheck.AvailLimit < input.Amount {
		return nil, fmt.Errorf("таны кредит гүйлгээний дүнд хүрэхгүй байна")
	}

	loanAccountID, err := a.client.Loan(int(input.Amount), "Зээл", int(input.CustomerID))
	if err != nil {
		return nil, fmt.Errorf("зээл авахад алдаа гарлаа: %w", err)
	}

	return &types.InvoiceResult{
		BankInvoiceID: fmt.Sprintf("%v", loanAccountID),
		IsPaid:        true,
		Raw:           loanAccountID,
	}, nil
}

func (a *BalcCreditAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("balc adapter not configured")
	}

	return &types.CheckInvoiceResult{
		IsPaid: true,
	}, nil
}
