package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.00
	expected := 6.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
