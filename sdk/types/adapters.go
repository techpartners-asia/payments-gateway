package types

type (
	QpayAdapter struct {
		Username    string
		Password    string
		Endpoint    string
		Callback    string
		InvoiceCode string
		MerchantID  string
	}

	TokipayAdapter struct {
		Endpoint      string
		APIKey        string
		IMAPIKey      string
		Authorization string
		MerchantID    string
		SuccessURL    string
		FailureURL    string
		AppSchemaIOS  string
	}

	StorePayAdapter struct {
		AppUserName string
		AppPassword string
		Username    string
		Password    string
		AuthUrl     string
		BaseUrl     string
		StoreId     string
		CallbackUrl string
	}

	SocialPayAdapter struct {
		Terminal string
		Secret   string
		Endpoint string
	}
	SimpleAdapter struct {
		UserName    string
		Password    string
		BaseUrl     string
		CallbackUrl string
	}
	PocketAdapter struct {
		Merchant      string
		ClientID      string
		ClientSecret  string
		Environment   string
		TerminalIDRaw int64
	}
	MonpayAdapter struct {
		Endpoint  string
		Username  string
		AccountID string
		Callback  string
	}
	GolomtAdapter struct {
		BaseURL     string
		Secret      string
		BearerToken string
	}
	BalcAdapter struct {
		Endpoint string
		Token    string
	}
)
