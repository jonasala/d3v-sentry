package d3vsentry

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)

//Setup configura a conexão com o Sentry
func Setup(dsn, environment string) error {
	return sentry.Init(sentry.ClientOptions{
		Dsn:         dsn,
		Environment: environment,
	})
}

//SentryRecovery é um middleware que captura panics e envia para o sentry através da sentry-go.
func SentryRecovery(repanic, waitForDelivery bool) func(handler http.Handler) http.Handler {
	return sentryhttp.New(sentryhttp.Options{
		Repanic:         repanic,
		WaitForDelivery: waitForDelivery,
	}).Handle
}

//SentryCapture envia err para o Sentry
func SentryCapture(err error) *sentry.EventID {
	id := sentry.CaptureException(err)
	sentry.Flush(time.Second * 5)
	return id
}
