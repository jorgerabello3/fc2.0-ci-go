package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 20)

	if total != 35 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 35)
	}
}
