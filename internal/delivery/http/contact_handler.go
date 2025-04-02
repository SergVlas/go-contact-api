package http

import (
	"app1/internal/domain"
	"app1/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ContactHandler struct {
	usecase  *usecase.ContactUsecase
	validate *validator.Validate
}

func NewContactHandler(uc *usecase.ContactUsecase) *ContactHandler {
	return &ContactHandler{
		usecase:  uc,
		validate: validator.New(),
	}
}

func (h *ContactHandler) Add(w http.ResponseWriter, r *http.Request) {
	var contact domain.Contact

	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(contact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.usecase.AddContact(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}
