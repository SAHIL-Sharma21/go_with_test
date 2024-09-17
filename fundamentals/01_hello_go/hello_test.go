package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello Golang, with Testing"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// func TestHelloUser(t *testing.T) {
// 	got := HelloUser("Sahil")
// 	want := "Hello, Sahil"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHelloUser(t *testing.T) {
	t.Run("saying Hello to user", func(t *testing.T) {
		got := HelloUser("Sahil")
		want := "Hello, Sahil"
		assertCorrectMessage(t, got, want)
	})

	//what if empty arg passed
	t.Run("Say Hello, User when an empty string is supplied", func(t *testing.T) {
		got := HelloUser("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

// making test case redable and refactoring
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreetUser(t *testing.T) {
	t.Run("in spanish", func(t *testing.T) {
		got := GreetUser("Sahil", "Spanish")
		want := "Hola, Sahil"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := GreetUser("Avinash", "French")
		want := "Bonjour, Avinash"

		assertCorrectMessage(t, got, want)
	})
}
