package utils

import protocol "github.com/tliron/glsp/protocol_3_16"

// UIntToInt converts LSP protocol.UInteger to int and apply a padding of 1
// because LSP based their positions on 0 but CUE on 1.
func UIntToInt(v protocol.UInteger) int {
	return int(v) + 1
}

// IntToUInt converts int position into protocol.UInteger with a padding of 1
// because LSP based their positions on 0 but CUE on 1.
func IntToUInt(v int) protocol.UInteger {
	return protocol.UInteger(v) - 1
}
