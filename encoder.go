package cbor

type Encoder struct {
	reader DataReader
}

func NewEncoder(buffer []byte) Encoder {
	return Encoder{
		reader: NewDataReader(buffer),
	}
}

// check whether any errors have occurred
func (e *Encoder) CheckError() error {
	return e.reader.CheckError()
}

// cbor ok
func (e *Encoder) WriteNil() {
	_ = e.reader.SetUint8(TypeNull)
}

// cbor ok
func (e *Encoder) WriteBool(value bool) {
	if value {
		_ = e.reader.SetUint8(TypeBoolTrue)
	} else {
		_ = e.reader.SetUint8(TypeBoolFalse)
	}
}

// cbor ok
func (e *Encoder) WriteInt8(value int8) {
	if value >= 0 {
		e.WriteUint8(uint8(value))
	} else {
		n := uint8(-1 - value)
		if n >= TypeU8ShortMin && n <= TypeU8ShortMax {
			_ = e.reader.SetUint8(TypeMajorSigned | n)
		} else {
			_ = e.reader.SetUint8(TypeI8)
			_ = e.reader.SetUint8(n)
		}
	}
}

// cbor ok
func (e *Encoder) WriteInt16(value int16) {
	if value >= 0 {
		e.WriteUint16(uint16(value))
	} else {
		n := uint16(-1 - value)
		if n <= TypeU8ShortMax {
			_ = e.reader.SetUint8(TypeMajorSigned | uint8(n))
		} else if n >= TypeU8 && n <= 0xff {
			_ = e.reader.SetUint8(TypeI8)
			_ = e.reader.SetUint8(uint8(n))
		} else {
			_ = e.reader.SetUint8(TypeI16)
			_ = e.reader.SetUint16(n)
		}
	}
}

// cbor ok
func (e *Encoder) WriteInt32(value int32) {
	if value >= 0 {
		e.WriteUint32(uint32(value))
	} else {
		n := uint32(-1 - value)
		if n <= TypeU8ShortMax {
			_ = e.reader.SetUint8(TypeMajorSigned | uint8(n))
		} else if n >= TypeU8 && n <= 0xff {
			_ = e.reader.SetUint8(TypeI8)
			_ = e.reader.SetUint8(uint8(n))
		} else if n >= 0x100 && n <= 0xffff {
			_ = e.reader.SetUint8(TypeI16)
			_ = e.reader.SetUint16(uint16(n))
		} else {
			_ = e.reader.SetUint8(TypeI32)
			_ = e.reader.SetUint32(n)
		}
	}
}

func (e *Encoder) WriteInt64(value int64) {
	if value >= 0 {
		e.WriteUint64(uint64(value))
	} else {
		n := uint64(-1 - value)
		if n <= TypeU8ShortMax {
			_ = e.reader.SetUint8(TypeMajorSigned | uint8(n))
		} else if n >= TypeU8 && n <= 0xff {
			_ = e.reader.SetUint8(TypeI8)
			_ = e.reader.SetUint8(uint8(n))
		} else if n >= 0x100 && n <= 0xffff {
			_ = e.reader.SetUint8(TypeI16)
			_ = e.reader.SetUint16(uint16(n))
		} else if n <= 0xffffffff {
			_ = e.reader.SetUint8(TypeI32)
			_ = e.reader.SetUint32(uint32(n))
		} else {
			_ = e.reader.SetUint8(TypeI64)
			_ = e.reader.SetUint64(n)
		}
	}
}

func (e *Encoder) WriteUint8(value uint8) {
	if value <= TypeU8ShortMax {
		_ = e.reader.SetUint8(value)
	} else {
		_ = e.reader.SetUint8(TypeU8)
		_ = e.reader.SetUint8(value)
	}
}

func (e *Encoder) WriteUint16(value uint16) {
	if value <= TypeU8ShortMax {
		_ = e.reader.SetUint8(uint8(value))
	} else if value <= 0xff {
		_ = e.reader.SetUint8(TypeU8)
		_ = e.reader.SetUint8(uint8(value))
	} else {
		_ = e.reader.SetUint8(TypeU16)
		_ = e.reader.SetUint16(value)
	}
}

func (e *Encoder) WriteUint32(value uint32) {
	if value <= TypeU8ShortMax {
		_ = e.reader.SetUint8(uint8(value))
	} else if value <= 0xff {
		_ = e.reader.SetUint8(TypeU8)
		_ = e.reader.SetUint8(uint8(value))
	} else if value <= 0xffff {
		_ = e.reader.SetUint8(TypeU16)
		_ = e.reader.SetUint16(uint16(value))
	} else {
		_ = e.reader.SetUint8(TypeU32)
		_ = e.reader.SetUint32(value)
	}
}

// cbor ok
func (e *Encoder) WriteUint64(value uint64) {
	if value <= TypeU8ShortMax {
		_ = e.reader.SetUint8(uint8(value))
	} else if value <= 0xff {
		_ = e.reader.SetUint8(TypeU8)
		_ = e.reader.SetUint8(uint8(value))
	} else if value <= 0xffff {
		_ = e.reader.SetUint8(TypeU16)
		_ = e.reader.SetUint16(uint16(value))
	} else if value <= 0xffffffff {
		_ = e.reader.SetUint8(TypeU32)
		_ = e.reader.SetUint32(uint32(value))
	} else {
		_ = e.reader.SetUint8(TypeU64)
		_ = e.reader.SetUint64(value)
	}
}

func (e *Encoder) WriteFloat32(value float32) {
	_ = e.reader.SetUint8(TypeF32)
	_ = e.reader.SetFloat32(value)
}

func (e *Encoder) WriteFloat64(value float64) {
	_ = e.reader.SetUint8(TypeF64)
	_ = e.reader.SetFloat64(value)
}

func (e *Encoder) writeTypeLength(t uint8, x uint64) {
	if x <= TypeU8ShortMax {
		_ = e.reader.SetUint8(t | uint8(x))
	} else if x < 0xff {
		_ = e.reader.SetUint8(t | 24)
		_ = e.reader.SetUint8(uint8(x))
	} else if x < 0xffff {
		_ = e.reader.SetUint8(t | 25)
		_ = e.reader.SetUint16(uint16(x))
	} else if x < 0xffffffff {
		_ = e.reader.SetUint8(t | 26)
		_ = e.reader.SetUint32(uint32(x))
	} else {
		_ = e.reader.SetUint8(t | 27)
		_ = e.reader.SetUint64(x)
	}
}

func (e *Encoder) WriteString(value string) {
	valueBytes := []byte(value)
	e.writeTypeLength(TypeMajorText, uint64(len(valueBytes)))
	_ = e.reader.SetBytes(valueBytes)
}

func (e *Encoder) WriteByteArray(value []byte) {
	e.writeTypeLength(TypeMajorBytes, uint64(len(value)))
	_ = e.reader.SetBytes(value)
}

func (e *Encoder) WriteArraySize(length uint32) {
	e.writeTypeLength(TypeMajorArray, uint64(length))
}

func (e *Encoder) WriteMapSize(length uint32) {
	e.writeTypeLength(TypeMajorMap, uint64(length))
}
