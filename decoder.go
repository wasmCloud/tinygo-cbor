package cbor

import (
	"math"
	"strconv"
)

type ReadError struct {
	message string
}

func (r ReadError) Error() string {
	return r.message
}

func NewReadError(s string) ReadError {
	return ReadError{message: s}
}

type Decoder struct {
	reader DataReader
}

func NewDecoder(buffer []byte) Decoder {
	return Decoder{
		reader: NewDataReader(buffer),
	}
}

// private function - exposed for debugging
func (d *Decoder) Pos() uint32 {
	return d.reader.byteOffset
}

func (d *Decoder) IsNextNil() (bool, error) {
	prefix, err := d.reader.PeekUint8()
	if err != nil {
		return false, err
	}
	if prefix == TypeNull {
		err = d.reader.Discard(1)
		return true, err
	}
	return false, nil
}

func (d *Decoder) ReadNull() (bool, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return false, err
	}
	if prefix == TypeNull {
		return true, nil
	}
	return false, ReadError{"bad value for null"}
}

func (d *Decoder) ReadBool() (bool, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return false, err
	} else if prefix == TypeBoolFalse {
		return false, nil
	} else if prefix == TypeBoolTrue {
		return true, nil
	}
	return false, NewReadError("bad value for bool")
}

func (d *Decoder) ReadInt8() (int8, error) {
	v, err := d.reader.GetUint8()
	if err != nil {
		return 0, ReadError{"int8"}
	}
	if v < TypeU8ShortMax {
		return int8(v), nil
	} else if v >= TypeMajorSigned && v <= TypeI8ShortMax {
		return -1 - (int8(v) - 0x20), nil
	}
	switch v {
	case TypeU8:
		n, err := d.reader.GetUint8()
		if err != nil {
			return 0, err
		}
		return int8(n), nil
	case TypeU16:
		n, err := d.reader.GetUint16()
		if err != nil || n > 0xff {
			return 0, ReadError{"invalid int8"}
		}
		return int8(n), nil
	case TypeU32:
		n, err := d.reader.GetUint32()
		if err != nil || n > 0xff {
			return 0, ReadError{"invalid int8"}
		}
		return int8(n), nil
	case TypeU64:
		n, err := d.reader.GetUint64()
		if err != nil || n > 0xff {
			return 0, ReadError{"invalid int8"}
		}
		return int8(n), nil
	case TypeI8:
		n, err := d.reader.GetUint8()
		if err != nil {
			return 0, err
		}
		return -1 - int8(n), nil
	case TypeI16:
		n, err := d.reader.GetUint16()
		if err != nil || n > 0xff {
			return 0, ReadError{"invalid int8"}
		}
		return int8(-1 - int8(n)), nil
	case TypeI32:
		n, err := d.reader.GetUint32()
		if err != nil || n > 0xff {
			return 0, ReadError{"invalid int8"}
		}
		return int8(-1 - int8(n)), nil
	case TypeI64:
		n, err := d.reader.GetUint64()
		if err != nil || n > 0xff {
			return 0, ReadError{"invalid int8"}
		}
		return int8(-1 - int8(n)), nil
	default:
		return 0, ReadError{"invalid int8"}
	}
}

func (d *Decoder) ReadInt16() (int16, error) {
	v, err := d.reader.GetUint8()
	if err != nil {
		return 0, ReadError{"int8"}
	}
	if v < TypeU8ShortMax {
		return int16(v), nil
	} else if v >= TypeMajorSigned && v <= TypeI8ShortMax {
		return -1 - (int16(v) - 0x20), nil
	}
	switch v {
	case TypeU8:
		n, err := d.reader.GetUint8()
		if err != nil {
			return 0, err
		}
		return int16(n), nil
	case TypeU16:
		n, err := d.reader.GetUint16()
		if err != nil || n > 0xffff {
			return 0, ReadError{"invalid int16"}
		}
		return int16(n), nil
	case TypeU32:
		n, err := d.reader.GetUint32()
		if err != nil || n > 0xffff {
			return 0, ReadError{"invalid int16"}
		}
		return int16(n), nil
	case TypeU64:
		n, err := d.reader.GetUint64()
		if err != nil || n > 0xffff {
			return 0, ReadError{"invalid int16"}
		}
		return int16(n), nil
	case TypeI8:
		n, err := d.reader.GetUint8()
		if err != nil {
			return 0, ReadError{"invalid int16"}
		}
		return int16(-1 - int16(n)), nil
	case TypeI16:
		n, err := d.reader.GetUint16()
		if err != nil {
			return 0, ReadError{"invalid int16"}
		}
		return -1 - int16(n), nil
	case TypeI32:
		n, err := d.reader.GetUint32()
		if err != nil || n > 0xffff {
			return 0, ReadError{"invalid int16"}
		}
		return int16(-1 - int32(n)), nil
	case TypeI64:
		n, err := d.reader.GetUint64()
		if err != nil || n > 0xffff {
			return 0, ReadError{"invalid int16"}
		}
		return int16(-1 - int16(n)), nil
	default:
		return 0, ReadError{"invalid int8"}
	}
}

func (d *Decoder) ReadInt32() (int32, error) {
	v, err := d.reader.GetUint8()
	if err != nil {
		return 0, ReadError{"int8"}
	}
	if v < TypeU8ShortMax {
		return int32(v), nil
	} else if v >= TypeMajorSigned && v <= TypeI8ShortMax {
		return -1 - (int32(v) - 0x20), nil
	}
	switch v {
	case TypeU8:
		n, err := d.reader.GetUint8()
		if err != nil {
			return 0, err
		}
		return int32(n), nil
	case TypeU16:
		n, err := d.reader.GetUint16()
		if err != nil {
			return 0, ReadError{"invalid uint32"}
		}
		return int32(n), nil
	case TypeU32:
		n, err := d.reader.GetUint32()
		if err != nil {
			return 0, ReadError{"invalid uint32"}
		}
		return int32(n), nil
	case TypeU64:
		n, err := d.reader.GetUint64()
		if err != nil || n > 0xffffffff {
			return 0, ReadError{"invalid uint32"}
		}
		return int32(n), nil
	case TypeI8:
		n, err := d.reader.GetUint8()
		if err != nil {
			return 0, err
		}
		return int32(-1 - int32(n)), nil
	case TypeI16:
		n, err := d.reader.GetUint16()
		if err != nil {
			return 0, ReadError{"invalid uint32"}
		}
		return int32(-1 - int32(n)), nil
	case TypeI32:
		n, err := d.reader.GetUint32()
		if err != nil {
			return 0, ReadError{"invalid uint32"}
		}
		return int32(-1 - int32(n)), nil
	case TypeI64:
		n, err := d.reader.GetUint64()
		if err != nil || n > 0xffffffff {
			return 0, ReadError{"invalid uint32"}
		}
		return int32(-1 - int32(n)), nil
	default:
		return 0, ReadError{"invalid uint32"}
	}
}

func (d *Decoder) ReadUint64() (uint64, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return 0, err
	}
	if prefix <= TypeU8ShortMax {
		return uint64(prefix), nil
	}
	switch prefix {
	case TypeU8:
		n, err_ := d.reader.GetUint8()
		if err_ != nil {
			return 0, err_
		}
		return uint64(n), nil
	case TypeU16:
		n, err_ := d.reader.GetUint16()
		return uint64(n), err_
	case TypeU32:
		n, err_ := d.reader.GetUint32()
		return uint64(n), err_
	case TypeU64:
		n, err_ := d.reader.GetUint64()
		return n, err_
	default:
		return 0, ReadError{"bad token for uint64"}
	}
}

func (d *Decoder) unsigned(prefix uint8) (uint64, error) {
	if prefix >= TypeU8ShortMin && prefix <= TypeU8ShortMax {
		return uint64(prefix), nil
	}
	switch prefix {
	case TypeU8:
		n, err_ := d.reader.GetUint8()
		return uint64(n), err_
	case TypeU16:
		n, err_ := d.reader.GetUint16()
		return uint64(n), err_
	case TypeU32:
		n, err_ := d.reader.GetUint32()
		return uint64(n), err_
	case TypeU64:
		n, err_ := d.reader.GetUint64()
		return n, err_
	default:
		return 0, ReadError{"bad token for uint64"}
	}
}

func (d *Decoder) ReadInt64() (int64, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return 0, err
	}
	if prefix >= TypeU8ShortMin && prefix <= TypeU8ShortMax {
		return int64(prefix), nil
	}
	if prefix >= TypeI8ShortMin && prefix <= TypeI8ShortMax {
		return int64(-1 - int8(prefix-TypeI8ShortMin)), nil
	}
	switch prefix {
	case TypeU8:
		n, err_ := d.reader.GetUint8()
		return int64(n), err_
	case TypeU16:
		n, err_ := d.reader.GetUint16()
		return int64(n), err_
	case TypeU32:
		n, err_ := d.reader.GetUint32()
		return int64(n), err_
	case TypeU64:
		n, err_ := d.reader.GetUint64()
		return int64(n), err_
	case TypeI8:
		n, err_ := d.reader.GetUint8()
		return -1 - int64(n), err_
	case TypeI16:
		n, err_ := d.reader.GetUint16()
		return -1 - int64(n), err_
	case TypeI32:
		n, err_ := d.reader.GetUint32()
		return -1 - int64(n), err_
	case TypeI64:
		n, err_ := d.reader.GetUint64()
		return -1 - int64(n), err_
	default:
		return 0, ReadError{"bad prefix for int64"}
	}
}

func (d *Decoder) ReadUint8() (uint8, error) {
	v, err := d.ReadUint64()
	if err != nil {
		return 0, err
	}
	if v <= 0xff {
		return uint8(v), nil
	}
	return 0, NewReadError(
		"integer overflow: value = " +
			strconv.FormatUint(v, 16) +
			"; bits = 8")
}

func (d *Decoder) ReadUint16() (uint16, error) {
	v, err := d.ReadUint64()
	if err != nil {
		return 0, err
	}
	if v <= math.MaxUint16 {
		return uint16(v), nil
	}
	return 0, ReadError{
		"integer overflow: value = " +
			strconv.FormatUint(v, 16) +
			"; bits = 16",
	}
}

func (d *Decoder) ReadUint32() (uint32, error) {
	v, err := d.ReadUint64()
	if err != nil {
		return 0, err
	}
	if v <= math.MaxUint32 {
		return uint32(v), nil
	}
	return 0, ReadError{
		"integer overflow: value = " +
			strconv.FormatUint(v, 16) +
			"; bits = 32",
	}
}

func (d *Decoder) ReadFloat32() (float32, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return 0, err
	}
	if prefix == TypeF32 {
		return d.reader.GetFloat32()
	}
	if prefix == TypeF16 {
		return 0, ReadError{"f16 not supported"}
	}
	return 0, ReadError{"bad prefix for float32"}
}

func (d *Decoder) ReadFloat64() (float64, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return 0, err
	}
	if prefix == TypeF32 {
		f, err_ := d.reader.GetFloat32()
		return float64(f), err_
	}
	if prefix == TypeF64 {
		return d.reader.GetFloat64()
	}
	return 0, ReadError{"bad prefix for float64"}
}

// Read string of defined length
func (d *Decoder) ReadString() (string, error) {
	strLen, err := d.readStringLength()
	if err != nil {
		return "", err
	}
	strBytes, err := d.reader.GetBytes(uint32(strLen))
	if err != nil {
		return "", err
	}
	return string(strBytes), nil
}

func (d *Decoder) readStringLength() (uint32, error) {
	b, err := d.reader.GetUint8()
	if err != nil {
		return 0, err
	}
	if TypeOf(b) != TypeMajorText || InfoOf(b) == 0x1f {
		return 0, ReadError{"expected string length"}
	}
	strLen, err := d.unsigned(InfoOf(b))
	if err != nil {
		return 0, ReadError{"expected string length"}
	}
	if strLen > 0xffffffff {
		return 0, ReadError{"string too long"}
	}
	return uint32(strLen), nil
}

func (d *Decoder) ReadByteArray() ([]byte, error) {
	binLen, err := d.readBinLength()
	if err != nil {
		return nil, err
	}
	binBytes, err := d.reader.GetBytes(binLen)
	if err != nil {
		return nil, err
	}
	return binBytes, nil
}

func (d *Decoder) readBinLength() (uint32, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return 0, err
	}
	if TypeOf(prefix) != TypeMajorBytes || InfoOf(prefix) == 31 {
		return 0, ReadError{"expected byte array (definite length)"}
	}
	n, err := d.unsigned(InfoOf(prefix))
	if err != nil {
		return 0, err
	}
	if n >= 0xffffffff {
		return 0, ReadError{"byte array too long"}
	}
	return uint32(n), nil
}

// For arrays of defined length, second value in return tuple should be false.
func (d *Decoder) ReadArraySize() (uint32, bool, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return 0, false, err
	}
	if TypeOf(prefix) != TypeMajorArray {
		return 0, false, ReadError{"expected array"}
	}
	b := InfoOf(prefix)
	if b == 31 {
		// indefinite array
		return 0, true, nil
	}
	arrLen, err := d.unsigned(b)
	if err != nil {
		return 0, false, err
	}
	if arrLen >= 0xffffffff {
		return 0, false, ReadError{"array too long"}
	}
	return uint32(arrLen), false, nil
}

func (d *Decoder) ReadMapSize() (uint32, bool, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return 0, false, err
	}
	if TypeOf(prefix) != TypeMajorMap {
		return 0, false, ReadError{"expected map"}
	}
	b := InfoOf(prefix)
	if b == 31 {
		// indefinite map
		return 0, true, nil
	}
	mapLen, err := d.unsigned(b)
	if err != nil {
		return 0, false, err
	}
	if mapLen >= 0xffffffff {
		return 0, false, ReadError{"map too large"}
	}
	return uint32(mapLen), false, nil
}

func (d *Decoder) ReadTag() (uint64, error) {
	prefix, err := d.reader.GetUint8()
	if err != nil {
		return 0, err
	}
	if TypeOf(prefix) != TypeMajorTagged {
		return 0, ReadError{"expected tag"}
	}
	return d.unsigned(InfoOf(prefix))
}

// subtract without underflow
func saturating_sub(a uint64, b uint64) uint64 {
	if b >= a {
		return 0
	}
	return a - b
}

func (d *Decoder) Skip() error {
	var nrounds uint64 = 1
	var irounds uint64 = 0

	for nrounds > 0 || irounds > 0 {
		peek, err := d.reader.PeekUint8()
		if err != nil {
			return err
		}
		if peek >= TypeU8ShortMin && peek <= TypeU64 {
			if _, err = d.ReadUint64(); err != nil {
				return err
			}
		} else if peek >= TypeMajorSigned && peek <= TypeI64 {
			if _, err = d.ReadInt64(); err != nil {
				return err
			}
		} else if peek >= TypeMajorBytes && peek <= TypeBytesIndef {
			n, err := d.readBinLength()
			if err != nil {
				// this will return error for indefinite byte len
				return err
			}
			if err = d.reader.Discard(n); err != nil {
				return err
			}
		} else if peek >= TypeMajorText && peek < TypeMajorArray {
			n, err := d.readStringLength()
			if err != nil {
				// this will return error for indefinite byte len
				return err
			}
			if err = d.reader.Discard(n); err != nil {
				return err
			}
		} else if peek >= TypeMajorArray && peek <= TypeArrayIndef {
			n, indef, err := d.ReadArraySize()
			if err != nil {
				return err
			}
			// if indef {
			//	return ReadError{"Indefinite arrays not supported"}
			// }
			if !indef {
				nrounds += uint64(n)
			} else if nrounds < 2 {
				irounds += 1
			} else {
				return ReadError{"no arrays of indefinite length inside regular arrays or maps"}
			}
		} else if peek >= TypeMajorMap && peek <= TypeMapIndef {
			n, indef, err := d.ReadMapSize()
			if err != nil {
				return err
			}
			if !indef {
				nrounds += uint64(n) * 2
			} else if nrounds < 2 {
				irounds += 1
			} else {
				return ReadError{"no maps of indefinite length inside regular arrays or maps"}
			}
		} else if peek >= TypeMajorTagged && peek <= TypeTagMax {
			n, err := d.ReadUint8()
			if err != nil {
				return err
			}
			_, err = d.unsigned(InfoOf(n))
			if err != nil {
				return err
			}
			continue
		} else if peek >= TypeMajorSimple && peek <= 0xfb {
			n, err := d.ReadUint8()
			if err != nil {
				return err
			}
			_, err = d.unsigned(InfoOf(n))
			if err != nil {
				return err
			}
		} else if peek == TypeBreak {
			_, err := d.ReadUint8()
			if err != nil {
				return err
			}
			irounds = saturating_sub(irounds, 1)
		} else {
			return ReadError{message: "unknown tag " + strconv.Itoa(int(peek)) + " @" + strconv.Itoa(int(d.reader.byteOffset))}
		}
		nrounds = saturating_sub(nrounds, 1)
	}
	return nil
}
