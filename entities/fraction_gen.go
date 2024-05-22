package entities

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/daoleno/uniswap-sdk-core/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Fraction) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Numerator")
			return
		}
		z.Numerator = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.Numerator))
			if err != nil {
				err = msgp.WrapError(err, "Numerator")
				return
			}
			z.Numerator = msgpencode.DecodeInt(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Denominator")
			return
		}
		z.Denominator = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.Denominator))
			if err != nil {
				err = msgp.WrapError(err, "Denominator")
				return
			}
			z.Denominator = msgpencode.DecodeInt(zb0003)
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Fraction) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	if z.Numerator == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Numerator))
		if err != nil {
			err = msgp.WrapError(err, "Numerator")
			return
		}
	}
	if z.Denominator == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Denominator))
		if err != nil {
			err = msgp.WrapError(err, "Denominator")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Fraction) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	if z.Numerator == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Numerator))
	}
	if z.Denominator == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Denominator))
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Fraction) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Numerator = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Numerator))
			if err != nil {
				err = msgp.WrapError(err, "Numerator")
				return
			}
			z.Numerator = msgpencode.DecodeInt(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Denominator = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Denominator))
			if err != nil {
				err = msgp.WrapError(err, "Denominator")
				return
			}
			z.Denominator = msgpencode.DecodeInt(zb0003)
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Fraction) Msgsize() (s int) {
	s = 1
	if z.Numerator == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Numerator))
	}
	if z.Denominator == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Denominator))
	}
	return
}