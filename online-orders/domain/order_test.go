package domain_test

import (
	"slices"
	"testing"

	"github.com/iamseki/online-orders/domain"
)

func TestFreteGratisLabel(t *testing.T) {
	payment := &domain.Payment{Value: 2000}
	product := &domain.Product{Category: domain.Infantil, Value: 2000}
	order := domain.NewOrder(payment, product)

	if !slices.Contains(order.Labels, domain.FreteGratis) {
		t.Errorf(`frete-gratis label not applied, expect [frete-gratis], but got: %v`, order.Labels)
	}
}

func TestEletrodomesticoLabel(t *testing.T) {
	payment := &domain.Payment{Value: 10}
	product := &domain.Product{Category: domain.Eletrodomestico, Value: 10}
	order := domain.NewOrder(payment, product)

	if !slices.Contains(order.Labels, domain.Fragil) {
		t.Errorf(`frete-gratis label not applied, expect [frágil], but got: %v`, order.Labels)
	}
}

func TestInfantilLabel(t *testing.T) {
	payment := &domain.Payment{Value: 100}
	product := &domain.Product{Category: domain.Infantil, Value: 100}
	order := domain.NewOrder(payment, product)

	if !slices.Contains(order.Labels, domain.Presente) {
		t.Errorf(`frete-gratis label not applied, expect [presente], but got: %v`, order.Labels)
	}
}

func TestBoletoDiscount(t *testing.T) {
	payment := &domain.Payment{Value: 100, Method: domain.Boleto}
	product := &domain.Product{Category: domain.Infantil, Value: 100}
	order := domain.NewOrder(payment, product)

	expectedPaymentValue := 90

	if expectedPaymentValue != order.Payment.Value {
		t.Errorf(`discount on Boleto payment method must be 10 percent, expect: %v, but got: %v`, expectedPaymentValue, order.Payment.Value)
	}
}
