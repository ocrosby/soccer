package common

import (
	"errors"
)

var ErrInvalidGender = errors.New("invalid gender")

type Gender int

const (
	Male Gender = iota
	Female
	UnspecifiedGender
)

func (g Gender) String() string {
	return [...]string{"Male", "Female", "Unspecified"}[g]
}
