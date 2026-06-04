package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, store.Fetch())
	}
}

type Store interface {
	Fetch() string
	Cancel()
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "Hello, world"
	server := Server(&SpyStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}

	t.Run("Tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello, world"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtxx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtxx)
	})
}
