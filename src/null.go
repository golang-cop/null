package Null

type Interface interface{}
type data struct{}

func New() Interface {
	return &data{}
}
