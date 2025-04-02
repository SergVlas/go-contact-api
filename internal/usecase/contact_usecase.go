package usecase

import (
	"app1/internal/domain"
)

type ContactUsecase struct {
	contactRepo domain.ContactRepository
}

func NewContactUsecase(repo domain.ContactRepository) *ContactUsecase {
	return &ContactUsecase{
		contactRepo: repo,
	}
}

func (uc *ContactUsecase) AddContact(contact *domain.Contact) error {
	// TODO валидация
	return uc.contactRepo.Add(contact)
}
