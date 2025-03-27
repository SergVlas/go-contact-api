package http

import (
	"app1/internal/config"
	//"app1/internal/delivery/http"
	"net/http"
)

func NewHandler(cfg *config.Config) (http.Handler, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page"))
	})

	//mux.HandleFunc("/contact", contacts.HandleAddContact)
	//mux.HandleFunc("/contact/list", contacts.HandleListContacts)

	return mux, nil
}
