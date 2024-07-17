package common

import (
	"encoding/json"
	"errors"
	"strings"
)

var ErrInvalidDivision = errors.New("invalid division")

type Division int

const (
	DI = iota
	DII
	DIII
	NAIA
	NJCAA
	UnspecifiedDivision
	TestDivision // TestDivision is used for testing purposes
)

func (d Division) String() string {
	return [...]string{"DI", "DII", "DIII", "NAIA", "NJCAA", "Unspecified"}[d]
}

func (d Division) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func StringToDivision(s string) (Division, error) {
	s = strings.TrimSpace(s)

	if s == "" {
		return UnspecifiedDivision, nil
	}

	s = strings.ToLower(s)

	switch s {
	case "di":
		return DI, nil
	case "dii":
		return DII, nil
	case "diii":
		return DIII, nil
	case "naia":
		return NAIA, nil
	case "njcaa":
		return NJCAA, nil
	default:
		return 0, ErrInvalidDivision
	}
}
