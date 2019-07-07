package http

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/sirnarsh/gotelltherabbit/rabbitmq"
	"github.com/sirnarsh/gotelltherabbit/readconf"
)

// StartServer starts http server at port 8080
func StartServer() {

	conf := readconf.GetGeneral()
	h2rConfig := readconf.GetH2R()
	handler := http.NewServeMux()

	handler.HandleFunc("/exchanges/", func(w http.ResponseWriter, r *http.Request) {

		exchangeName := strings.Replace(r.URL.Path, "/exchanges/", "", 1)

		// Check if exchange is while listed
		isExchangeAllowed := false
		for _, allowedExchange := range h2rConfig.AllowedExchanges {
			if allowedExchange == "*" {
				isExchangeAllowed = true
			}
			if allowedExchange == exchangeName {
				isExchangeAllowed = true
			}
		}
		if !isExchangeAllowed {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "text/plain")
			io.WriteString(w, "You are not allowed to send this exchange")
		}
		// Read body data
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			io.WriteString(w, "Couldn't read request")
		}

		// Send rabbitmq message
		// @todo implement error handling (right now connection is dropped on error while sending)
		rabbitmq.Send(exchangeName, bodyBytes)
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "Successfully sent to rabbitmq")
	})

	handler.HandleFunc("/exchanges", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Please add exchange name in url")
	})

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "GoTellTheRabbit server running")
	})

	err := http.ListenAndServe(conf.ServerBind, handler)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
