package Calculadora

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// Definimos los valores de prueba
	num1 := 1
	num2 := 1
	expectedResult := 2

	// Llamamos a la función Add con los valores de prueba
	result := Add(num1, num2)

	// Comparamos el resultado devuelto con el resultado esperado
	if result != expectedResult {
		t.Errorf("Add(%d, %d) = %d; expected %d", num1, num2, result, expectedResult)
	}
}
