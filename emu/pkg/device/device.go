package device

type D8Device interface {
	Read(uint16) uint8
	Write(uint16, uint8) uint8
	ReadEnabled(uint16) bool
	WriteEnabled(uint16) bool
}