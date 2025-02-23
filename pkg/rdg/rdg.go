package rdg

import "github.com/google/uuid"

const publicIDLength = 8

func GenerateID() uuid.UUID {
	return uuid.New()
}

func GeneratePublicID() string {
	return uuid.NewString()[:publicIDLength]
}
