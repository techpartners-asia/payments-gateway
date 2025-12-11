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

func NewTokiPayAdapter(input types.TokipayAdapter) *TokiPayAdapter {
	return &TokiPayAdapter{client: tokipay.New(input.Endpoint, input.APIKey, input.IMAPIKey, input.Authorization, input.MerchantID, input.SuccessURL, input.FailureURL, input.AppSchemaIOS)}
}

func (a *TokiPayAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("tokipay adapter not configured")
	}
	res, err := a.client.PaymentSentUser(tokipay.TokipayPaymentInput{
		OrderId:     input.UID,
		Amount:      int64(input.Amount),
		PhoneNo:     input.Phone,
		CountryCode: "+976",
		Notes:       input.Note,
	})
	if err != nil {
		return nil, err
	}
	return &types.InvoiceResult{
		BankInvoiceID: input.UID,
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *TokiPayAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("tokipay adapter not configured")
	}

	res, err := a.client.PaymentStatus(input.UID)
	if err != nil {
		return nil, err
	}

	return &types.CheckInvoiceResult{
		IsPaid: res.Data.Status == "COMPLETED",
	}, nil
}
