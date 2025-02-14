package uuid

import (
	"fmt"
	"github.com/google/uuid"
)

func StringToUUID(s string) (uuid.UUID, error) {
	var id uuid.UUID

	id, err := uuid.Parse(s)

	if err != nil {
		return id, fmt.Errorf("error while parsing UUID field to string")
	}

	return id, nil
}
