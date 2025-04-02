package memory

import (
	"sync"

	"app1/internal/domain"

	"github.com/google/uuid"
)

type ContactRepo struct {
	contacts map[string]*domain.Contact
	mu       sync.RWMutex
}

func NewContactRepo() *ContactRepo {
	return &ContactRepo{
		contacts: make(map[string]*domain.Contact),
	}
}

func (r *ContactRepo) Add(contact *domain.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for {
		contact.ID = uuid.New().String()
		if _, exists := r.contacts[contact.ID]; !exists {
			break
		}
	}

	r.contacts[contact.ID] = contact
	return nil
}
