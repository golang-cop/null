package Null

// Interface is the Null-Object primitive. Its defining property is that
// IsNull() answers true, so callers can branch on a value's null-ness through
// the interface rather than comparing against a bare nil.
type Interface interface {
	IsNull() bool
}

type data struct{}

func New() Interface {
	return &data{}
}

func (d data) IsNull() bool {
	return true
}
