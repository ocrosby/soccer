package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidGender = errors.New("invalid gender")

// Gender represents the gender of an individual or group.
// swagger:enum
type Gender int

const (
	// Male represents the male gender.
	Male Gender = iota
	// Female represents the female gender.
	Female
	// UnspecifiedGender represents an unspecified gender.
	UnspecifiedGender
)

func (g Gender) String() string {
	return [...]string{"Male", "Female", "Unspecified"}[g]
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

func StringToGender(s string) (Gender, error) {
	s = strings.TrimSpace(s) // Trim leading and trailing spaces
	s = strings.ToLower(s)   // Convert s to lowercase

	// Return UnspecifiedGender if s is empty
	if len(s) == 0 {
		return UnspecifiedGender, nil
	}

	switch s {
	case "male":
		return Male, nil
	case "female":
		return Female, nil
	case "unspecified":
		return UnspecifiedGender, nil
	default:
		return 0, fmt.Errorf("invalid gender: %s", s)
	}
}
