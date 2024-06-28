package main

import "testing"

func TestShouldSumIncorrectly(t *testing.T){
		test:= sum(2,2)
		result:= 7

		if test != result {
		 t.Error("Valor esperado da soma", result, "Valor retornado", test)
	}
}

func TestShouldSumCorrectly(t *testing.T){
	test:= sum(2,2)
	result:= 4

	if test != result {
		t.Error("Valor esperado da soma", result, "Valor retornado", test)
}
}

func TestShoulSubtractIncorrectly(t *testing.T){
	test:= subtract(2,2)
	result:= 1

	if test != result {
		t.Error("Valor esperado da subtração", result, "Valor retornado", test)
}
}

func TestShouldSubtractCorrectly(t *testing.T){
test:= subtract(2,2)
result:= 0

		if test != result {
			t.Error("Valor esperado da soma subtração", result, "Valor retornado", test)
		}
}

func TestShouldMultipyIncorrectly(t *testing.T){
	test:= multiplication(2,2)
	result:= 6

	if test != result {
		t.Error("Valor esperado da multiplicação", result, "Valor retornado", test)
}
}

func TestShouldMultiplyCorrectly(t *testing.T){
test:= multiplication(2,2)
result:= 4

		if test != result {
			t.Error("Valor esperado da soma multiplicação", result, "Valor retornado", test)
		}
}

func TestShouldDivisionIncorrectly(t *testing.T){
	test:= division(100,10)
	result:= 1

	if test != float32(result) {
		t.Error("Valor esperado da divisão", result, "Valor retornado", test)
}
}

func TestShouldDivisionCorrectly(t *testing.T){
test:= division(100,10)
result:= 10

		if test != float32(result) {
			t.Error("Valor esperado da soma multiplicação", result, "Valor retornado", test)
		}
}