package print

import (
	"strings"
	"testing"
)

func TestPrintStr(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Fatalf("got '%q' want '%q' ", got, want)
		}

	}
	t.Run("日本人说话", func(t *testing.T) {
		got := PrintStr("せかい", strings.ToUpper("japanese"))
		want := "こんにちは～せかい1"
		assertCorrectMessage(t, got, want)
	})
	t.Run("中国人说话", func(t *testing.T) {
		got := PrintStr("世界", strings.ToUpper("chinese"))
		want := "你好！世界1"
		assertCorrectMessage(t, got, want)
	})
	t.Run("其他老外说话", func(t *testing.T) {
		got := PrintStr("world", strings.ToUpper("english"))
		want := "Hello world1"
		assertCorrectMessage(t, got, want)
	})

}
