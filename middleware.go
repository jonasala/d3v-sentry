package d3v_sentry

import (
	"net/http"
	raven "github.com/getsentry/raven-go"
)

//SentryRecovery é um middleware que captura panics e envia para o sentry através da raven-go. O parâmetro repanic determina se após o envio o panic será reenviado
//**IMPORTANTE** É necessário configurar o raven-go na inicialização do programa: https://docs.sentry.io/clients/go/integrations/
func SentryRecovery(repanic bool) func(handler *http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)) {
		defer func() {
			if rval := recover(); rval != nil {
				rvalString := fmt.Sprint(rval)

				var packet *raven.Packet

				packet = raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.GetOrNewStacktrace(err, 2, 3, nil)), raven.NewHttp(r))
				if err, ok := rval.(error); ok {
					packet = raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.GetOrNewStacktrace(err, 2, 3, nil)), raven.NewHttp(r))
				} else {
					packet = raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
				}
				raven.Capture(packet, nil)
				if (repanic) {
					panic(rval)
				} else {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}()
		handler.ServeHTTP(w, r)
	}
}