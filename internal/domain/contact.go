package domain

type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type ContactRepository interface {
	Add(contact *Contact) error
}
