package Null

type Interface interface{}
type data struct{}

func New() Interface {
	return &data{}
}

func (d data) IsNul() bool {
	return true
}
