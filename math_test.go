package main

import "testing"

func TestSoma(t *testing.T) {

	resultado := Soma(4, 10)

	if resultado != 15 {
		t.Errorf("Resultado = %d; esperado 15", resultado)
	}
}
