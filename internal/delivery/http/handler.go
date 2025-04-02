package http

import (
	"app1/internal/config"
	"net/http"
)

func NewHandler(cfg *config.Config, contactHandler *ContactHandler) (http.Handler, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page"))
	})

	mux.HandleFunc("/contact", contactHandler.Add)
	//mux.HandleFunc("/contact/list", contacts.HandleListContacts)

	return mux, nil
}
