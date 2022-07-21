package server

type Mode int

const (
	ModeProd Mode = iota
	ModeDev
)

func (m Mode) IsProd() bool {
	if m == ModeProd {
		return true
	}
	return false
}
