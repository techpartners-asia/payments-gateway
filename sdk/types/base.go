package types

type PaymentType string

const (
	PaymentTypeQPay     PaymentType = "qpay"
	PaymentTypeTokipay  PaymentType = "tokipay"
	PaymentTypeMonpay   PaymentType = "monpay"
	PaymentTypeGolomt   PaymentType = "golomt"
	PaymentTypeSocial   PaymentType = "socialpay"
	PaymentTypeStorePay PaymentType = "storepay"
	PaymentTypePocket   PaymentType = "pocket"
	PaymentTypeSimple   PaymentType = "simple"
	PaymentTypeBalc     PaymentType = "balc"
)

type (
	InvoiceInput struct {
		Amount      float64     // Amount
		UID         string      // payment uid or order uid
		Phone       string      // phone number
		CustomerID  uint        // customer id
		Note        string      // Note : description of the invoice
		CallbackURL string      // CallbackURL : callback url
		ReturnType  string      // ReturnType : return type
		Type        PaymentType // qpay , tokipay , monpay , golomt , socialpay , storepay , pocket , simple , balc
	}

	InvoiceResult struct {
		BankInvoiceID string     `json:"bank_invoice_id"`
		BankQRCode    string     `json:"bank_qr_code"`
		Deeplinks     []Deeplink `json:"deeplinks"`
		IsPaid        bool       `json:"is_paid"`
		Raw           any        `json:"raw"`
	}

	Deeplink struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Logo        string `json:"logo"`
	}

	CheckInvoiceInput struct {
		UID    string      `json:"uid"`
		Amount float64     `json:"amount"`
		Type   PaymentType `json:"type"`
	}

	CheckInvoiceResult struct {
		IsPaid bool   `json:"is_paid"`
		Msg    string `json:"msg"`
	}
)
