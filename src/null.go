package Null

// Interface is the Null-Object primitive. Its defining property is that
// IsNull() answers true, so callers can branch on a value's null-ness through
// the interface rather than comparing against a bare nil.
type Interface interface {
	IsNull() bool
}

type data struct{}

// theNull is the single, immutable Null-Object. A Null carries no state and
// IsNull() is always true, so every Null is interchangeable — one shared
// instance is semantically identical to a fresh one, and allocates nothing.
// (The Null-Object pattern is what makes this interning safe.)
var theNull = &data{}

func New() Interface {
	return theNull
}

func (d data) IsNull() bool {
	return true
}
