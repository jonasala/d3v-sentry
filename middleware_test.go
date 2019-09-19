package d3v_sentry_test

import (
	"net/http"

	"github.com/go-chi/chi"
	sentry "github.com/jonasala/d3v-sentry"
)

func ExampleSentryRecovery() {
	r := chi.NewRouter()

	// Apply the middleware to the router
	r.Use(sentry.SentryRecovery(true))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		panic("catched")
	})
}
