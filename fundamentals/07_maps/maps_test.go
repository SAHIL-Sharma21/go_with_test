package main

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "This is just a test"}

	got := Search(dict, "test")
	want := "This is just a test"

	assertStrings(t, got, want)

}

// 1st refactoring
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
