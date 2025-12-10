## Payments SDK (Go)

Unified gateway over multiple payment providers (QPay, Tokipay, Golomt Ecommerce, SocialPay, StorePay, Pocket, Simple, Balc; Monpay stubbed).

### Install

```
go get github.com/techpartners-asia/payments-gateway/sdk
```

### Quick Start

```go
import paymentssdk "github.com/techpartners-asia/payments-gateway/sdk"

cfg := paymentssdk.Config{
    Qpay: paymentssdk.QPayConfig{
        Username:    os.Getenv("QPAY_USERNAME"),
        Password:    os.Getenv("QPAY_PASSWORD"),
        Endpoint:    os.Getenv("QPAY_ENDPOINT"),
        Callback:    os.Getenv("QPAY_CALLBACK"),
        InvoiceCode: os.Getenv("QPAY_INVOICE_CODE"),
        MerchantID:  os.Getenv("QPAY_MERCHANT_ID"),
    },
    // fill other providers as needed...
}

gw, err := paymentssdk.New(cfg)
if err != nil { log.Fatal(err) }

res, err := gw.CreateInvoice(paymentssdk.InvoiceInput{
    Type:       paymentssdk.PaymentTypeQPay,
    Amount:     15000,
    PaymentUID: "order-123",
    IsOrg:      false,
    OrgRegNo:   "",
    CallbackURL: cfg.Qpay.Callback, // optional per provider
})
if err != nil { log.Fatal(err) }

fmt.Println("invoice:", res.BankInvoiceID, "qr:", res.BankQRCode, "deeplinks:", res.Deeplinks)
```

### Provider Notes (required fields)

- **QPay:** username, password, endpoint, callback, invoiceCode, merchantID.
- **Tokipay:** endpoint, apiKey, imApiKey, authorization, merchantID, successURL, failureURL, appSchemaIOS.
- **Golomt Ecommerce:** baseURL, secret, bearerToken, plus `CallbackURL` and optional `ReturnType` in `InvoiceInput` (`GET`|`POST`|`MOBILE`).
- **SocialPay:** terminal, secret, endpoint.
- **StorePay:** appUsername/appPassword, username/password, authURL, baseURL, storeID, callbackURL.
- **Pocket:** merchant, clientID, clientSecret, environment, terminalIDRaw (string, parsed to int64).
- **Simple:** username, password, baseURL, callbackURL; optional `ExpireMinutes` in `InvoiceInput` (default 20).
- **Balc:** endpoint, token; marks `IsPaid=true` on create.
- **Monpay:** create-invoice not implemented; use library QR helpers directly.

### Packages

- `sdk` (package `paymentssdk`): public entrypoints, config, gateway wiring.
- `sdk/types`: shared types (re-exported by `sdk/aliases.go`).
- `sdk/adapters`: per-provider adapters implementing `PaymentProvider`.

### Public API

- `New(cfg Config) (*Gateway, error)` – build SDK clients from config, return gateway.
- `NewGatewayFromSDK(*SDK) *Gateway` – if you already constructed clients.
- `Gateway.CreateInvoice(InvoiceInput) (*InvoiceResult, error)` – route by `PaymentType`.

### Model Types

- `PaymentType*` constants (qpay, tokipay, monpay, golomt, socialpay, storepay, pocket, simple, balc).
- `InvoiceInput` – unified request per payment.
- `InvoiceResult` – normalized response (invoice id, QR, deeplinks, raw payload, isPaid).

### Caveats

- Monpay adapter returns an error for create; use monpay package QR helpers instead.
- Ensure provider configs are complete; validation fails fast on missing required fields.
