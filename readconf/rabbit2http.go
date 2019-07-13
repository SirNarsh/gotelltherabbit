package readconf

// HeaderKeyValue Key value pairs of headers
type HeaderKeyValue struct {
	Key, Value string
}

// BindExchanges  Object for exchange binding
type BindExchanges struct {
	ExchangeName, HTTPURL string
	HTTPHeaders           []HeaderKeyValue
}

// Rabbit2http root config
type Rabbit2http struct {
	ListenQueue   string
	BindExchanges []BindExchanges
}
