package http_model

type Market struct {
	Symbols []Symbol
}

type Symbol struct {
	BaseAsset  string
	QuoteAsset string
	Status     string
	// Filters    []interface{}
}
