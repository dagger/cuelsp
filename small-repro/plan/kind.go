//go:generate stringer -type=Kind

package plan

type Kind uint8

const (
	File Kind = iota
	Directory
)
