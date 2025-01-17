package properties

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Journal) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "26934511":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentPosted")
					return
				}
				z.LogDocumentPosted = nil
			} else {
				if z.LogDocumentPosted == nil {
					z.LogDocumentPosted = new(bool)
				}
				*z.LogDocumentPosted, err = dc.ReadBool()
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentPosted")
					return
				}
			}
		case "26932611":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentPrinted")
					return
				}
				z.LogDocumentPrinted = nil
			} else {
				if z.LogDocumentPrinted == nil {
					z.LogDocumentPrinted = new(bool)
				}
				*z.LogDocumentPrinted, err = dc.ReadBool()
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentPrinted")
					return
				}
			}
		case "26934411":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentRouted")
					return
				}
				z.LogDocumentRouted = nil
			} else {
				if z.LogDocumentRouted == nil {
					z.LogDocumentRouted = new(bool)
				}
				*z.LogDocumentRouted, err = dc.ReadBool()
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentRouted")
					return
				}
			}
		case "26932711":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentSaved")
					return
				}
				z.LogDocumentSaved = nil
			} else {
				if z.LogDocumentSaved == nil {
					z.LogDocumentSaved = new(bool)
				}
				*z.LogDocumentSaved, err = dc.ReadBool()
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentSaved")
					return
				}
			}
		case "2693193":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogDuration")
					return
				}
				z.LogDuration = nil
			} else {
				if z.LogDuration == nil {
					z.LogDuration = new(int32)
				}
				*z.LogDuration, err = dc.ReadInt32()
				if err != nil {
					err = msgp.WrapError(err, "LogDuration")
					return
				}
			}
		case "26932064":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogEnd")
					return
				}
				z.LogEnd = nil
			} else {
				if z.LogEnd == nil {
					z.LogEnd = new(int64)
				}
				*z.LogEnd, err = dc.ReadInt64()
				if err != nil {
					err = msgp.WrapError(err, "LogEnd")
					return
				}
			}
		case "2693243":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogFlags")
					return
				}
				z.LogFlags = nil
			} else {
				if z.LogFlags == nil {
					z.LogFlags = new(int32)
				}
				*z.LogFlags, err = dc.ReadInt32()
				if err != nil {
					err = msgp.WrapError(err, "LogFlags")
					return
				}
			}
		case "26931864":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogStart")
					return
				}
				z.LogStart = nil
			} else {
				if z.LogStart == nil {
					z.LogStart = new(int64)
				}
				*z.LogStart, err = dc.ReadInt64()
				if err != nil {
					err = msgp.WrapError(err, "LogStart")
					return
				}
			}
		case "26931231":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogType")
					return
				}
				z.LogType = nil
			} else {
				if z.LogType == nil {
					z.LogType = new(string)
				}
				*z.LogType, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "LogType")
					return
				}
			}
		case "26934631":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "LogTypeDesc")
					return
				}
				z.LogTypeDesc = nil
			} else {
				if z.LogTypeDesc == nil {
					z.LogTypeDesc = new(string)
				}
				*z.LogTypeDesc, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "LogTypeDesc")
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Journal) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "26934511"
	err = en.Append(0x8a, 0xa8, 0x32, 0x36, 0x39, 0x33, 0x34, 0x35, 0x31, 0x31)
	if err != nil {
		return
	}
	if z.LogDocumentPosted == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBool(*z.LogDocumentPosted)
		if err != nil {
			err = msgp.WrapError(err, "LogDocumentPosted")
			return
		}
	}
	// write "26932611"
	err = en.Append(0xa8, 0x32, 0x36, 0x39, 0x33, 0x32, 0x36, 0x31, 0x31)
	if err != nil {
		return
	}
	if z.LogDocumentPrinted == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBool(*z.LogDocumentPrinted)
		if err != nil {
			err = msgp.WrapError(err, "LogDocumentPrinted")
			return
		}
	}
	// write "26934411"
	err = en.Append(0xa8, 0x32, 0x36, 0x39, 0x33, 0x34, 0x34, 0x31, 0x31)
	if err != nil {
		return
	}
	if z.LogDocumentRouted == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBool(*z.LogDocumentRouted)
		if err != nil {
			err = msgp.WrapError(err, "LogDocumentRouted")
			return
		}
	}
	// write "26932711"
	err = en.Append(0xa8, 0x32, 0x36, 0x39, 0x33, 0x32, 0x37, 0x31, 0x31)
	if err != nil {
		return
	}
	if z.LogDocumentSaved == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBool(*z.LogDocumentSaved)
		if err != nil {
			err = msgp.WrapError(err, "LogDocumentSaved")
			return
		}
	}
	// write "2693193"
	err = en.Append(0xa7, 0x32, 0x36, 0x39, 0x33, 0x31, 0x39, 0x33)
	if err != nil {
		return
	}
	if z.LogDuration == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteInt32(*z.LogDuration)
		if err != nil {
			err = msgp.WrapError(err, "LogDuration")
			return
		}
	}
	// write "26932064"
	err = en.Append(0xa8, 0x32, 0x36, 0x39, 0x33, 0x32, 0x30, 0x36, 0x34)
	if err != nil {
		return
	}
	if z.LogEnd == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteInt64(*z.LogEnd)
		if err != nil {
			err = msgp.WrapError(err, "LogEnd")
			return
		}
	}
	// write "2693243"
	err = en.Append(0xa7, 0x32, 0x36, 0x39, 0x33, 0x32, 0x34, 0x33)
	if err != nil {
		return
	}
	if z.LogFlags == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteInt32(*z.LogFlags)
		if err != nil {
			err = msgp.WrapError(err, "LogFlags")
			return
		}
	}
	// write "26931864"
	err = en.Append(0xa8, 0x32, 0x36, 0x39, 0x33, 0x31, 0x38, 0x36, 0x34)
	if err != nil {
		return
	}
	if z.LogStart == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteInt64(*z.LogStart)
		if err != nil {
			err = msgp.WrapError(err, "LogStart")
			return
		}
	}
	// write "26931231"
	err = en.Append(0xa8, 0x32, 0x36, 0x39, 0x33, 0x31, 0x32, 0x33, 0x31)
	if err != nil {
		return
	}
	if z.LogType == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteString(*z.LogType)
		if err != nil {
			err = msgp.WrapError(err, "LogType")
			return
		}
	}
	// write "26934631"
	err = en.Append(0xa8, 0x32, 0x36, 0x39, 0x33, 0x34, 0x36, 0x33, 0x31)
	if err != nil {
		return
	}
	if z.LogTypeDesc == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteString(*z.LogTypeDesc)
		if err != nil {
			err = msgp.WrapError(err, "LogTypeDesc")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Journal) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "26934511"
	o = append(o, 0x8a, 0xa8, 0x32, 0x36, 0x39, 0x33, 0x34, 0x35, 0x31, 0x31)
	if z.LogDocumentPosted == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBool(o, *z.LogDocumentPosted)
	}
	// string "26932611"
	o = append(o, 0xa8, 0x32, 0x36, 0x39, 0x33, 0x32, 0x36, 0x31, 0x31)
	if z.LogDocumentPrinted == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBool(o, *z.LogDocumentPrinted)
	}
	// string "26934411"
	o = append(o, 0xa8, 0x32, 0x36, 0x39, 0x33, 0x34, 0x34, 0x31, 0x31)
	if z.LogDocumentRouted == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBool(o, *z.LogDocumentRouted)
	}
	// string "26932711"
	o = append(o, 0xa8, 0x32, 0x36, 0x39, 0x33, 0x32, 0x37, 0x31, 0x31)
	if z.LogDocumentSaved == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBool(o, *z.LogDocumentSaved)
	}
	// string "2693193"
	o = append(o, 0xa7, 0x32, 0x36, 0x39, 0x33, 0x31, 0x39, 0x33)
	if z.LogDuration == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendInt32(o, *z.LogDuration)
	}
	// string "26932064"
	o = append(o, 0xa8, 0x32, 0x36, 0x39, 0x33, 0x32, 0x30, 0x36, 0x34)
	if z.LogEnd == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendInt64(o, *z.LogEnd)
	}
	// string "2693243"
	o = append(o, 0xa7, 0x32, 0x36, 0x39, 0x33, 0x32, 0x34, 0x33)
	if z.LogFlags == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendInt32(o, *z.LogFlags)
	}
	// string "26931864"
	o = append(o, 0xa8, 0x32, 0x36, 0x39, 0x33, 0x31, 0x38, 0x36, 0x34)
	if z.LogStart == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendInt64(o, *z.LogStart)
	}
	// string "26931231"
	o = append(o, 0xa8, 0x32, 0x36, 0x39, 0x33, 0x31, 0x32, 0x33, 0x31)
	if z.LogType == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendString(o, *z.LogType)
	}
	// string "26934631"
	o = append(o, 0xa8, 0x32, 0x36, 0x39, 0x33, 0x34, 0x36, 0x33, 0x31)
	if z.LogTypeDesc == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendString(o, *z.LogTypeDesc)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Journal) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "26934511":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogDocumentPosted = nil
			} else {
				if z.LogDocumentPosted == nil {
					z.LogDocumentPosted = new(bool)
				}
				*z.LogDocumentPosted, bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentPosted")
					return
				}
			}
		case "26932611":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogDocumentPrinted = nil
			} else {
				if z.LogDocumentPrinted == nil {
					z.LogDocumentPrinted = new(bool)
				}
				*z.LogDocumentPrinted, bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentPrinted")
					return
				}
			}
		case "26934411":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogDocumentRouted = nil
			} else {
				if z.LogDocumentRouted == nil {
					z.LogDocumentRouted = new(bool)
				}
				*z.LogDocumentRouted, bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentRouted")
					return
				}
			}
		case "26932711":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogDocumentSaved = nil
			} else {
				if z.LogDocumentSaved == nil {
					z.LogDocumentSaved = new(bool)
				}
				*z.LogDocumentSaved, bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogDocumentSaved")
					return
				}
			}
		case "2693193":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogDuration = nil
			} else {
				if z.LogDuration == nil {
					z.LogDuration = new(int32)
				}
				*z.LogDuration, bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogDuration")
					return
				}
			}
		case "26932064":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogEnd = nil
			} else {
				if z.LogEnd == nil {
					z.LogEnd = new(int64)
				}
				*z.LogEnd, bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogEnd")
					return
				}
			}
		case "2693243":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogFlags = nil
			} else {
				if z.LogFlags == nil {
					z.LogFlags = new(int32)
				}
				*z.LogFlags, bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogFlags")
					return
				}
			}
		case "26931864":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogStart = nil
			} else {
				if z.LogStart == nil {
					z.LogStart = new(int64)
				}
				*z.LogStart, bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogStart")
					return
				}
			}
		case "26931231":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogType = nil
			} else {
				if z.LogType == nil {
					z.LogType = new(string)
				}
				*z.LogType, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogType")
					return
				}
			}
		case "26934631":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LogTypeDesc = nil
			} else {
				if z.LogTypeDesc == nil {
					z.LogTypeDesc = new(string)
				}
				*z.LogTypeDesc, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogTypeDesc")
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Journal) Msgsize() (s int) {
	s = 1 + 9
	if z.LogDocumentPosted == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BoolSize
	}
	s += 9
	if z.LogDocumentPrinted == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BoolSize
	}
	s += 9
	if z.LogDocumentRouted == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BoolSize
	}
	s += 9
	if z.LogDocumentSaved == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BoolSize
	}
	s += 8
	if z.LogDuration == nil {
		s += msgp.NilSize
	} else {
		s += msgp.Int32Size
	}
	s += 9
	if z.LogEnd == nil {
		s += msgp.NilSize
	} else {
		s += msgp.Int64Size
	}
	s += 8
	if z.LogFlags == nil {
		s += msgp.NilSize
	} else {
		s += msgp.Int32Size
	}
	s += 9
	if z.LogStart == nil {
		s += msgp.NilSize
	} else {
		s += msgp.Int64Size
	}
	s += 9
	if z.LogType == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.LogType)
	}
	s += 9
	if z.LogTypeDesc == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.LogTypeDesc)
	}
	return
}
