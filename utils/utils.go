package utils
import (
	"github.com/google/uuid"
)
func GenerateUUID() string {
	return uuid.New().String()
}

func StringPtr(s string) *string {
	return &s
}