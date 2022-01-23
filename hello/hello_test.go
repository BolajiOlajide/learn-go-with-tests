package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello without an argument", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello without an argument in english", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello without an argument in spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello without an argument in french", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people without specifying a language", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in english", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello top people in french", func(t *testing.T) {
		got := Hello("Max", "French")
		want := "Bonjour, Max"
		assertCorrectMessage(t, got, want)
	})
}
