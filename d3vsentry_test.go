package d3vsentry

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	assert.Nil(t, Setup("https://d1965561cbbf4c398704b605725ee5bb@sentry.io/1740891", "test"))
}

func TestSentryRecovery(t *testing.T) {
	Setup("https://d1965561cbbf4c398704b605725ee5bb@sentry.io/1740891", "test")

	router := chi.NewRouter()
	router.Use(SentryRecovery(false, true))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		panic("d3v-sentry test")
	})

	ts := httptest.NewServer(router)

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Failed to capture panic. Expected 200 got %v", resp.StatusCode)
	}
}

func TestSentryCapture(t *testing.T) {
	Setup("https://d1965561cbbf4c398704b605725ee5bb@sentry.io/1740891", "test")
	assert.NotNil(t, SentryCapture(errors.New("d3v-sentry test capture")))
}
