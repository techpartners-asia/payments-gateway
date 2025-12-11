package sdkAdapters

import (
	"fmt"
	"strconv"

	"github.com/techpartners-asia/payments-gateway/sdk/types"

	"github.com/techpartners-asia/qpay-go/qpay_v2"
)

// QPayAdapter implements PaymentProvider for QPay.
type QPayAdapter struct {
	client qpay_v2.QPay
}

func NewQPayAdapter(input types.QpayAdapter) *QPayAdapter {
	if input.Username == "" || input.Password == "" || input.Endpoint == "" || input.InvoiceCode == "" || input.MerchantID == "" {
		return nil
	}
	return &QPayAdapter{client: qpay_v2.New(input.Username, input.Password, input.Endpoint, input.Callback, input.InvoiceCode, input.MerchantID)}
}

func (a *QPayAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("qpay adapter not configured")
	}
	// prefix := "personal"
	// if input.IsOrg && input.OrgRegNo != "" {
	// 	prefix = input.OrgRegNo
	// }
	qpayInput := qpay_v2.QPayCreateInvoiceInput{
		SenderCode:    input.UID,
		ReceiverCode:  input.UID,
		Description:   input.Note,
		Amount:        int64(input.Amount),
		CallbackParam: map[string]string{"uid": input.UID},
	}
	res, _, err := a.client.CreateInvoice(qpayInput)
	if err != nil {
		return nil, err
	}

	var deeplinks []types.Deeplink
	for _, v := range res.Urls {
		deeplinks = append(deeplinks, types.Deeplink{
			Name:        v.Name,
			Description: v.Description,
			Link:        v.Link,
			Logo:        v.Logo,
		})
	}

	return &types.InvoiceResult{
		BankInvoiceID: res.InvoiceID,
		BankQRCode:    res.QrText,
		Deeplinks:     deeplinks,
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *QPayAdapter) CheckInvoice(input types.CheckInvoiceInput) (*types.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("qpay adapter not configured")
	}

	res, _, err := a.client.CheckPayment(input.UID, 100, 1)
	if err != nil {
		return nil, err
	}

	amount := float64(0)
	for _, row := range res.Rows {
		if row.PaymentStatus == "PAID" {
			if _amount, err := strconv.ParseFloat(row.PaymentAmount, 64); err == nil {
				amount += _amount
			}
		}
	}

	return &types.CheckInvoiceResult{
		IsPaid: amount >= input.Amount,
	}, nil
}
