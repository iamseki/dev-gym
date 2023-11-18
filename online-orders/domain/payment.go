package domain

type PaymentMethod int64

const BOLETO_DISCOUNT = 0.1

const (
	Boleto PaymentMethod = iota
)

type Payment struct {
	Method PaymentMethod
	Value  int
}

func (p PaymentMethod) String() string {
	switch p {
	case Boleto:
		return "boleto"
	}
	return "unknown"
}

func (p *Payment) applyBoletoDiscount(productValue int) {
	p.Value = int(float64(productValue) - float64(productValue)*BOLETO_DISCOUNT)
}
