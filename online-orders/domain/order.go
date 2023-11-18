package domain

type OrderLabel int64

const (
	Fragil OrderLabel = iota
	FreteGratis
	Presente
)

func (ol OrderLabel) String() string {
	switch ol {
	case Fragil:
		return "frágil"
	case FreteGratis:
		return "frete-gratis"
	case Presente:
		return "presente"
	}
	return "unknown"
}

type Order struct {
	*Product
	*Payment
	Labels []OrderLabel
}

func NewOrder(payment *Payment, product *Product) *Order {
	order := &Order{
		Payment: payment,
		Product: product,
	}

	order.applyLabels()
	order.applyDiscounts()

	return order
}

func (o *Order) applyDiscounts() {
	switch o.Payment.Method {
	case Boleto:
		o.Payment.applyBoletoDiscount(o.Product.Value)
	}
}

func (o *Order) applyLabels() {
	if o.Payment.Value > 1000 {
		o.Labels = append(o.Labels, FreteGratis)
	}

	switch o.Product.Category {
	case Eletrodomestico:
		o.Labels = append(o.Labels, Fragil)
	case Infantil:
		o.Labels = append(o.Labels, Presente)
	}
}
