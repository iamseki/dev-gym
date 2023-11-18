package domain

type ProductCategory int64

const (
	Eletrodomestico ProductCategory = iota
	Infantil
)

type Product struct {
	Category ProductCategory
	Value    int
}

func (p ProductCategory) String() string {
	switch p {
	case Eletrodomestico:
		return "eletrodoméstico"
	case Infantil:
		return "infantil"
	}
	return "unknown"
}
