package json

type Generator interface {
	Generate() []byte
}
