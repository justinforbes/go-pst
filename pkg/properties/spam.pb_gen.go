package properties

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Spam) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "248353":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "JunkAddRecipientsToSafeSendersList")
					return
				}
				z.JunkAddRecipientsToSafeSendersList = nil
			} else {
				if z.JunkAddRecipientsToSafeSendersList == nil {
					z.JunkAddRecipientsToSafeSendersList = new(int32)
				}
				*z.JunkAddRecipientsToSafeSendersList, err = dc.ReadInt32()
				if err != nil {
					err = msgp.WrapError(err, "JunkAddRecipientsToSafeSendersList")
					return
				}
			}
		case "248323":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "JunkIncludeContacts")
					return
				}
				z.JunkIncludeContacts = nil
			} else {
				if z.JunkIncludeContacts == nil {
					z.JunkIncludeContacts = new(int32)
				}
				*z.JunkIncludeContacts, err = dc.ReadInt32()
				if err != nil {
					err = msgp.WrapError(err, "JunkIncludeContacts")
					return
				}
			}
		case "248343":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "JunkPermanentlyDelete")
					return
				}
				z.JunkPermanentlyDelete = nil
			} else {
				if z.JunkPermanentlyDelete == nil {
					z.JunkPermanentlyDelete = new(int32)
				}
				*z.JunkPermanentlyDelete, err = dc.ReadInt32()
				if err != nil {
					err = msgp.WrapError(err, "JunkPermanentlyDelete")
					return
				}
			}
		case "2483911":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "JunkPhishingEnableLinks")
					return
				}
				z.JunkPhishingEnableLinks = nil
			} else {
				if z.JunkPhishingEnableLinks == nil {
					z.JunkPhishingEnableLinks = new(bool)
				}
				*z.JunkPhishingEnableLinks, err = dc.ReadBool()
				if err != nil {
					err = msgp.WrapError(err, "JunkPhishingEnableLinks")
					return
				}
			}
		case "248333":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "JunkThreshold")
					return
				}
				z.JunkThreshold = nil
			} else {
				if z.JunkThreshold == nil {
					z.JunkThreshold = new(int32)
				}
				*z.JunkThreshold, err = dc.ReadInt32()
				if err != nil {
					err = msgp.WrapError(err, "JunkThreshold")
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
func (z *Spam) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "248353"
	err = en.Append(0x85, 0xa6, 0x32, 0x34, 0x38, 0x33, 0x35, 0x33)
	if err != nil {
		return
	}
	if z.JunkAddRecipientsToSafeSendersList == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteInt32(*z.JunkAddRecipientsToSafeSendersList)
		if err != nil {
			err = msgp.WrapError(err, "JunkAddRecipientsToSafeSendersList")
			return
		}
	}
	// write "248323"
	err = en.Append(0xa6, 0x32, 0x34, 0x38, 0x33, 0x32, 0x33)
	if err != nil {
		return
	}
	if z.JunkIncludeContacts == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteInt32(*z.JunkIncludeContacts)
		if err != nil {
			err = msgp.WrapError(err, "JunkIncludeContacts")
			return
		}
	}
	// write "248343"
	err = en.Append(0xa6, 0x32, 0x34, 0x38, 0x33, 0x34, 0x33)
	if err != nil {
		return
	}
	if z.JunkPermanentlyDelete == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteInt32(*z.JunkPermanentlyDelete)
		if err != nil {
			err = msgp.WrapError(err, "JunkPermanentlyDelete")
			return
		}
	}
	// write "2483911"
	err = en.Append(0xa7, 0x32, 0x34, 0x38, 0x33, 0x39, 0x31, 0x31)
	if err != nil {
		return
	}
	if z.JunkPhishingEnableLinks == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBool(*z.JunkPhishingEnableLinks)
		if err != nil {
			err = msgp.WrapError(err, "JunkPhishingEnableLinks")
			return
		}
	}
	// write "248333"
	err = en.Append(0xa6, 0x32, 0x34, 0x38, 0x33, 0x33, 0x33)
	if err != nil {
		return
	}
	if z.JunkThreshold == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteInt32(*z.JunkThreshold)
		if err != nil {
			err = msgp.WrapError(err, "JunkThreshold")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Spam) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "248353"
	o = append(o, 0x85, 0xa6, 0x32, 0x34, 0x38, 0x33, 0x35, 0x33)
	if z.JunkAddRecipientsToSafeSendersList == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendInt32(o, *z.JunkAddRecipientsToSafeSendersList)
	}
	// string "248323"
	o = append(o, 0xa6, 0x32, 0x34, 0x38, 0x33, 0x32, 0x33)
	if z.JunkIncludeContacts == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendInt32(o, *z.JunkIncludeContacts)
	}
	// string "248343"
	o = append(o, 0xa6, 0x32, 0x34, 0x38, 0x33, 0x34, 0x33)
	if z.JunkPermanentlyDelete == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendInt32(o, *z.JunkPermanentlyDelete)
	}
	// string "2483911"
	o = append(o, 0xa7, 0x32, 0x34, 0x38, 0x33, 0x39, 0x31, 0x31)
	if z.JunkPhishingEnableLinks == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBool(o, *z.JunkPhishingEnableLinks)
	}
	// string "248333"
	o = append(o, 0xa6, 0x32, 0x34, 0x38, 0x33, 0x33, 0x33)
	if z.JunkThreshold == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendInt32(o, *z.JunkThreshold)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Spam) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "248353":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.JunkAddRecipientsToSafeSendersList = nil
			} else {
				if z.JunkAddRecipientsToSafeSendersList == nil {
					z.JunkAddRecipientsToSafeSendersList = new(int32)
				}
				*z.JunkAddRecipientsToSafeSendersList, bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "JunkAddRecipientsToSafeSendersList")
					return
				}
			}
		case "248323":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.JunkIncludeContacts = nil
			} else {
				if z.JunkIncludeContacts == nil {
					z.JunkIncludeContacts = new(int32)
				}
				*z.JunkIncludeContacts, bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "JunkIncludeContacts")
					return
				}
			}
		case "248343":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.JunkPermanentlyDelete = nil
			} else {
				if z.JunkPermanentlyDelete == nil {
					z.JunkPermanentlyDelete = new(int32)
				}
				*z.JunkPermanentlyDelete, bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "JunkPermanentlyDelete")
					return
				}
			}
		case "2483911":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.JunkPhishingEnableLinks = nil
			} else {
				if z.JunkPhishingEnableLinks == nil {
					z.JunkPhishingEnableLinks = new(bool)
				}
				*z.JunkPhishingEnableLinks, bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "JunkPhishingEnableLinks")
					return
				}
			}
		case "248333":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.JunkThreshold = nil
			} else {
				if z.JunkThreshold == nil {
					z.JunkThreshold = new(int32)
				}
				*z.JunkThreshold, bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "JunkThreshold")
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
func (z *Spam) Msgsize() (s int) {
	s = 1 + 7
	if z.JunkAddRecipientsToSafeSendersList == nil {
		s += msgp.NilSize
	} else {
		s += msgp.Int32Size
	}
	s += 7
	if z.JunkIncludeContacts == nil {
		s += msgp.NilSize
	} else {
		s += msgp.Int32Size
	}
	s += 7
	if z.JunkPermanentlyDelete == nil {
		s += msgp.NilSize
	} else {
		s += msgp.Int32Size
	}
	s += 8
	if z.JunkPhishingEnableLinks == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BoolSize
	}
	s += 7
	if z.JunkThreshold == nil {
		s += msgp.NilSize
	} else {
		s += msgp.Int32Size
	}
	return
}
