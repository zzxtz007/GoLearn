package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// 1
	// slowURL := "http://www.facebook.com"
	// fastURL := "http://www.baidu.com"

	// want := fastURL
	// got := Racer(slowURL, fastURL)

	// if got != want {
	// 	t.Errorf("got '%s' ,want '%s'", got, want)
	// }
	// 2
	//
	// slowServer := makeDelayedServer(20 * time.Millisecond)

	// fastServer := makeDelayedServer(0 * time.Millisecond)

	// //defer 结束时调用
	// defer slowServer.Close()
	// defer fastServer.Close()

	// slowURL := slowServer.URL
	// fastURL := fastServer.URL

	// want := fastURL
	// got, _ := Racer(slowURL, fastURL, tenSecondTimeout)

	// if got != want {
	// 	t.Errorf("got '%s', want '%s'", got, want)
	// }

	// t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
	// 	serverA := makeDelayedServer(11 * time.Second)
	// 	serverB := makeDelayedServer(12 * time.Second)

	// 	defer serverA.Close()
	// 	defer serverB.Close()

	// 	_, err := Racer(serverA.URL, serverB.URL, tenSecondTimeout)

	// 	if err == nil {
	// 		t.Error("expected an error but didn't get one")
	// 	}
	// })
	// 3
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
