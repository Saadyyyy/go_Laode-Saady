package calculator

import "testing"

func TestAddtion(t *testing.T) {
	if Addtion(10, 20) != 30 {
		t.Error("Expected output 10 + 20 to equel 30")
	}

	if Addtion(-10, -20) != -30 {
		t.Error("Expected output -10 + -20 to equel -30")
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(10, 20) != -10 {
		t.Error("Expected  10 (-) 20 to equal -10")
	}

	if Subtraction(-5, -10) != 5 {
		t.Error("Expected  -5 (-) -10 to equal 5")
	}
}

func TestDivision(t *testing.T) {
	if Division(10, 20) != 0.5 {
		t.Error("Expected  10 (/) 20 to equel 0.5")
	}

	if Division(-5, -10) != 0.5 {
		t.Error("Expected  -5 (/) -10 to equel 0.5")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(10, 20) != 200 {
		t.Error("Expected  10 (*) 20 to equel 200")
	}

	if Multiplication(-5, -10) != 50 {
		t.Error("Expected  -5 (*) -10 to equel 50")
	}
}
