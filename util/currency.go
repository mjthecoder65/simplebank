package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	KRW = "KRW"
	TZS = "TZS"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, KRW, TZS:
		return true
	}
	return false
}
