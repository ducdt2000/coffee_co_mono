package payment

type Means string

const (
	MEANS_CARD     = "card"
	MEANS_CASH     = "cash"
	MEANS_COFFEBUX = "coffeebux"
)

type CardDetails struct {
	cardToken string
}
