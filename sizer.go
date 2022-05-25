package cbor

type Sizer struct {
	length uint32
}

func NewSizer() Sizer {
	return Sizer{}
}

// check whether any errors have occurred
func (s *Sizer) CheckError() error {
	return nil
}

func (s *Sizer) Len() uint32 {
	return s.length
}

func (s *Sizer) WriteNil() {
	s.length++
}

func (s *Sizer) writeTypeLength(t uint8, x uint64) {
	if x < TypeU8ShortMax {
		s.length++
	} else if x <= 0xff {
		s.length += 2
	} else if x <= 0xffff {
		s.length += 3
	} else if x <= 0xffffffff {
		s.length += 5
	} else {
		s.length += 9
	}
}

func (s *Sizer) WriteString(value string) {
	buf := []byte(value)
	s.writeTypeLength(TypeMajorText, uint64(len(buf)))
	s.length += uint32(len(buf))
}

func (s *Sizer) WriteBool(value bool) {
	s.length++
}

func (s *Sizer) WriteArraySize(length uint32) {
	s.writeTypeLength(TypeMajorArray, uint64(length))
}

func (s *Sizer) WriteByteArray(value []byte) {
	s.writeTypeLength(TypeMajorArray, uint64(len(value)))
	s.length += uint32(len(value))
}

func (s *Sizer) WriteMapSize(length uint32) {
	s.writeTypeLength(TypeMajorMap, uint64(length))
}

func (s *Sizer) WriteInt8(value int8) {
	s.WriteInt64(int64(value))
}
func (s *Sizer) WriteInt16(value int16) {
	s.WriteInt64(int64(value))
}
func (s *Sizer) WriteInt32(value int32) {
	s.WriteInt64(int64(value))
}
func (s *Sizer) WriteInt64(value int64) {
	if value > 0 {
		s.WriteUint64(uint64(value))
	} else {
		n := uint64(-1 - value)
		if n <= TypeU8ShortMax {
			s.length += 1
		} else if n >= TypeU8 && n <= 0xff {
			s.length += 2
		} else if n >= 0x100 && n <= 0xffff {
			s.length += 3
		} else if n <= 0xffffffff {
			s.length += 5
		} else {
			s.length += 9
		}
	}
}

func (s *Sizer) WriteUint8(value uint8) {
	s.WriteUint64(uint64(value))
}
func (s *Sizer) WriteUint16(value uint16) {
	s.WriteUint64(uint64(value))
}
func (s *Sizer) WriteUint32(value uint32) {
	s.WriteUint64(uint64(value))
}
func (s *Sizer) WriteUint64(value uint64) {
	if value <= TypeU8ShortMax {
		s.length++
	} else if value <= 0xff {
		s.length += 2
	} else if value <= 0xffff {
		s.length += 3
	} else if value <= 0xffffffff {
		s.length += 5
	} else {
		s.length += 9
	}
}

func (s *Sizer) WriteFloat32(value float32) {
	s.length += 5
}
func (s *Sizer) WriteFloat64(value float64) {
	s.length += 9
}
