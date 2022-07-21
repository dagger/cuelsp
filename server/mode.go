package server

type Mode int

const (
	ModeProd Mode = iota
	ModeDev
)

func (m Mode) IsProd() bool {
	switch m {
	case ModeProd:
		return true
	default:
		return false
	}
}
