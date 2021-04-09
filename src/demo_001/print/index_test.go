package print

import "testing"

func TestPrintStr(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := PrintStr("123")
		want := "123"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := PrintStr("")
		want := "Hello Wr2old"
		assertCorrectMessage(t, got, want)
	})

}
