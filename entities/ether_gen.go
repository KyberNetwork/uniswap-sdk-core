package entities

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *etherInner) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "BaseCurrency")
			return
		}
		z.BaseCurrency = nil
	} else {
		if z.BaseCurrency == nil {
			z.BaseCurrency = new(BaseCurrency)
		}
		err = z.BaseCurrency.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "BaseCurrency")
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *etherInner) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	if z.BaseCurrency == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.BaseCurrency.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "BaseCurrency")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *etherInner) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	if z.BaseCurrency == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.BaseCurrency.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "BaseCurrency")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *etherInner) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.BaseCurrency = nil
	} else {
		if z.BaseCurrency == nil {
			z.BaseCurrency = new(BaseCurrency)
		}
		bts, err = z.BaseCurrency.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "BaseCurrency")
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *etherInner) Msgsize() (s int) {
	s = 1
	if z.BaseCurrency == nil {
		s += msgp.NilSize
	} else {
		s += z.BaseCurrency.Msgsize()
	}
	return
}
