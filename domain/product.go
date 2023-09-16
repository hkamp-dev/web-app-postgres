package domain

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
}
