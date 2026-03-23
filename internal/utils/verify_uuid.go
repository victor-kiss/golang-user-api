package utils

import (
	"errors"

	"github.com/google/uuid"
)

func ValidateUUID(UUID string) error {

	if UUID == "" {
		return errors.New("UUID cant be empty")
	}

	_, err := uuid.Parse(UUID)

	if err != nil {
		return errors.New("UUID invalid format")
	}
	return nil
}
