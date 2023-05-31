package coffee_co

import "github.com/google/uuid"

type CoffeeLover struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	EmailAddress string
}
