package main

import (
	"testing"
)

func TestHello(t *testing.T){

	t.Run("Greeting people", func(t *testing.T){	
		actual := Hello("Alan","")
		expected := "Hello, Alan"

		if actual != expected {
			t.Errorf("actual %q expected %q",actual,expected)
		} 
	})

	t.Run("say 'Hello, World when an empty string is supplied", func(t *testing.T){
		actual := Hello("","")
		expected := "Hello, World"

		if actual != expected{
			t.Errorf("actual %q expected %q",actual, expected)
		}
	})
	t.Run("in Spanish", func(t *testing.T){
		actual := Hello("Alan", "spanish")
		expected := "Hola, Alan"

	
		if actual != expected {
			t.Errorf("actual %q expected %q",actual,expected)	
		}
	})
	t.Run("In French", func(t *testing.T){
		actual := Hello("Alan", "french")
		expected := "Bonjour, Alan"
		
		if actual != expected{
			t.Errorf("actual %q expected %q", actual, expected)
		}
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string){
	t.Helper()
	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}

