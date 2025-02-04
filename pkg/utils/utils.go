package utils

import "github.com/google/uuid"

func GetRandID(length int) string {
	return uuid.NewString()[:length]
}
