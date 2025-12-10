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

func NewQPayAdapter(client qpay_v2.QPay) *QPayAdapter {
	return &QPayAdapter{client: client}
}

func (a *QPayAdapter) CreateInvoice(input types.InvoiceInput) (*types.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("qpay adapter not configured")
	}
	prefix := "personal"
	if input.IsOrg && input.OrgRegNo != "" {
		prefix = input.OrgRegNo
	}
	qpayInput := qpay_v2.QPayCreateInvoiceInput{
		SenderCode:    fmt.Sprintf("%s-%s", prefix, input.PaymentUID),
		ReceiverCode:  prefix,
		Description:   "Захиалга",
		Amount:        int64(input.Amount),
		CallbackParam: map[string]string{"payment_uid": input.PaymentUID},
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

	res, _, err := a.client.CheckPayment(input.PaymentUID, 100, 1)
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
