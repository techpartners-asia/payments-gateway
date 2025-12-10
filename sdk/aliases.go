package paymentssdk

import "github.com/techpartners-asia/payments-gateway/sdk/types"

// Re-export common types for convenient imports.
type (
	Config          = types.Config
	QPayConfig      = types.QPayConfig
	TokipayConfig   = types.TokipayConfig
	MonpayConfig    = types.MonpayConfig
	GolomtConfig    = types.GolomtConfig
	SocialPayConfig = types.SocialPayConfig
	StorePayConfig  = types.StorePayConfig
	PocketConfig    = types.PocketConfig
	SimpleConfig    = types.SimpleConfig
	BalcAPIConfig   = types.BalcAPIConfig

	InvoiceInput    = types.InvoiceInput
	InvoiceResult   = types.InvoiceResult
	PaymentType     = types.PaymentType
	PaymentProvider = types.PaymentProvider
)

const (
	PaymentTypeQPay     = types.PaymentTypeQPay
	PaymentTypeTokipay  = types.PaymentTypeTokipay
	PaymentTypeMonpay   = types.PaymentTypeMonpay
	PaymentTypeGolomt   = types.PaymentTypeGolomt
	PaymentTypeSocial   = types.PaymentTypeSocial
	PaymentTypeStorePay = types.PaymentTypeStorePay
	PaymentTypePocket   = types.PaymentTypePocket
	PaymentTypeSimple   = types.PaymentTypeSimple
	PaymentTypeBalc     = types.PaymentTypeBalc
)
