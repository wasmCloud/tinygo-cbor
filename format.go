package cbor

// Get major type if the byte (high 3 bits)
func TypeOf(b uint8) uint8 {
	return b & 0xe0
}

// Get additional info of the byte (low 5 bits)
func InfoOf(b uint8) uint8 {
	return b & 0x1f
}

const (
	TypeU8ShortMin     = 0x00
	TypeU8ShortMax     = 0x17
	TypeU8             = 0x18
	TypeU16            = 0x19
	TypeU32            = 0x1a
	TypeU64            = 0x1b
	TypeI8ShortMin     = 0x20
	TypeI8ShortMax     = 0x37
	TypeI8             = 0x38
	TypeI16            = 0x39
	TypeI32            = 0x3a
	TypeI64            = 0x3b
	TypeBytesIndef     = 0x5f
	TypeArrayIndef     = 0x9f
	TypeMapIndef       = 0xbf
	TypeTagMax         = 0xdb
	TypeBoolFalse      = 0xf4
	TypeBoolTrue       = 0xf5
	TypeNull           = 0xf6
	TypeUndefined      = 0xf7
	TypeF16            = 0xf9
	TypeF32            = 0xfa
	TypeF64            = 0xfb
	TypeBreak          = 0xff
	TypeMajorUnsigned  = 0x00 // major type : high 3 bits
	TypeMajorSigned    = 0x20 // major type : high 3 bits
	TypeMajorBytes     = 0x40 // major type : high 3 bits
	TypeMajorText      = 0x60 // major type : high 3 bits
	TypeMajorArray     = 0x80 // major type : high 3 bits
	TypeMajorMap       = 0xa0 // major type : high 3 bits
	TypeMajorTagged    = 0xc0 // major type : high 3 bits
	TypeMajorSimple    = 0xe0 // major type : high 3 bits
)
