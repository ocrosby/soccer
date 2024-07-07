package common

import "errors"

var ErrInvalidDivision = errors.New("invalid division")

type Division int

const (
	DI = iota
	DII
	DIII
	NAIA
	NJCAA
	UnspecifiedDivision
)

func (d Division) String() string {
	return [...]string{"DI", "DII", "DIII", "NAIA", "NJCAA", "Unspecified"}[d]
}
