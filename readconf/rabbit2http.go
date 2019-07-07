package readconf

type BindExchanges struct {
	ExchangeName, HttpURL string
	HttpHeaders           []string
}

type Rabbit2http struct {
	ListenQueue   string
	BindExchanges []BindExchanges
}
