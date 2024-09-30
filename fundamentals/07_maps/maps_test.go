package main

import "testing"

// func TestSearch(t *testing.T) {
// 	dict := map[string]string{"test": "This is just a test"}

// 	got := Search(dict, "test")
// 	want := "This is just a test"

// 	assertStrings(t, got, want)

// }

// 1st refactoring
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// func TestSearch(t *testing.T) {
// 	dict := Dictinory{"test": "this is just a test"}

// 	got := dict.Search("test")
// 	want := "this is just a test"

// 	assertStrings(t, got, want)
// }

// test case when there is key present or not
func TestSearch(t *testing.T) {
	dict := Dictinory{"test": "this is just a test"}
	t.Run("known key", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dict.Search("unknown")
		if err == nil {
			t.Fatal("expected to get the error.")
		}
		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

// func TestAdd(t *testing.T) {
// 	dict := Dictinory{}
// 	dict.Add("test", "this is just a test")

// 	want := "this is just a test"
// 	got, err := dict.Search("test")

// 	if err != nil {
// 		t.Fatal("Should find added word", err)
// 	}
// 	assertStrings(t, got, want)
// }

// refactoring the test add
// func TestAdd(t *testing.T) {
// 	dict := Dictinory{}
// 	word := "test"
// 	defination := "this is just a test"

// 	dict.Add(word, defination)
// 	assertDefination(t, dict, word, defination)
// }

func assertDefination(t testing.TB, dict Dictinory, word, defination string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, defination)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictinory{}
		word := "test"
		defination := "this is just a test"

		err := dict.Add(word, defination)

		assertError(t, err, nil)
		assertDefination(t, dict, word, defination)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dict := Dictinory{word: defination}

		err := dict.Add(word, "new test")
		assertError(t, err, nil)
		assertDefination(t, dict, word, defination)
	})
}

// func TestUpdate(t *testing.T) {
// 	word := "test"
// 	defination := "this is just a test"
// 	dict := Dictinory{word: defination}
// 	newDefination := "new defination"

// 	dict.Update(word, newDefination)
// 	assertDefination(t, dict, word, newDefination)
// }

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dict := Dictinory{word: defination}
		newDefination := "new defination"

		err := dict.Update(word, newDefination)

		assertError(t, err, nil)
		assertDefination(t, dict, word, newDefination)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dict := Dictinory{}

		err := dict.Update(word, defination)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

// function delete
func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictinory{word: "test defination"}

	dict.Delete(word)

	_, err := dict.Search(word)
	assertError(t, err, ErrNotFound)
}
