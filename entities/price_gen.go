package entities

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Price) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Fraction")
			return
		}
		z.Fraction = nil
	} else {
		if z.Fraction == nil {
			z.Fraction = new(Fraction)
		}
		err = z.Fraction.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "Fraction")
			return
		}
	}
	{
		var zb0002 []byte
		zb0002, err = dc.ReadBytes(EncodeCurrency(z.BaseCurrency))
		if err != nil {
			err = msgp.WrapError(err, "BaseCurrency")
			return
		}
		z.BaseCurrency = DecodeCurrency(zb0002)
	}
	{
		var zb0003 []byte
		zb0003, err = dc.ReadBytes(EncodeCurrency(z.QuoteCurrency))
		if err != nil {
			err = msgp.WrapError(err, "QuoteCurrency")
			return
		}
		z.QuoteCurrency = DecodeCurrency(zb0003)
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Scalar")
			return
		}
		z.Scalar = nil
	} else {
		if z.Scalar == nil {
			z.Scalar = new(Fraction)
		}
		err = z.Scalar.DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "Scalar")
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Price) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	if z.Fraction == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Fraction.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Fraction")
			return
		}
	}
	err = en.WriteBytes(EncodeCurrency(z.BaseCurrency))
	if err != nil {
		err = msgp.WrapError(err, "BaseCurrency")
		return
	}
	err = en.WriteBytes(EncodeCurrency(z.QuoteCurrency))
	if err != nil {
		err = msgp.WrapError(err, "QuoteCurrency")
		return
	}
	if z.Scalar == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Scalar.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Scalar")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Price) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	if z.Fraction == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Fraction.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Fraction")
			return
		}
	}
	o = msgp.AppendBytes(o, EncodeCurrency(z.BaseCurrency))
	o = msgp.AppendBytes(o, EncodeCurrency(z.QuoteCurrency))
	if z.Scalar == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Scalar.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Scalar")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Price) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Fraction = nil
	} else {
		if z.Fraction == nil {
			z.Fraction = new(Fraction)
		}
		bts, err = z.Fraction.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "Fraction")
			return
		}
	}
	{
		var zb0002 []byte
		zb0002, bts, err = msgp.ReadBytesBytes(bts, EncodeCurrency(z.BaseCurrency))
		if err != nil {
			err = msgp.WrapError(err, "BaseCurrency")
			return
		}
		z.BaseCurrency = DecodeCurrency(zb0002)
	}
	{
		var zb0003 []byte
		zb0003, bts, err = msgp.ReadBytesBytes(bts, EncodeCurrency(z.QuoteCurrency))
		if err != nil {
			err = msgp.WrapError(err, "QuoteCurrency")
			return
		}
		z.QuoteCurrency = DecodeCurrency(zb0003)
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Scalar = nil
	} else {
		if z.Scalar == nil {
			z.Scalar = new(Fraction)
		}
		bts, err = z.Scalar.UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "Scalar")
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Price) Msgsize() (s int) {
	s = 1
	if z.Fraction == nil {
		s += msgp.NilSize
	} else {
		s += z.Fraction.Msgsize()
	}
	s += msgp.BytesPrefixSize + len(EncodeCurrency(z.BaseCurrency)) + msgp.BytesPrefixSize + len(EncodeCurrency(z.QuoteCurrency))
	if z.Scalar == nil {
		s += msgp.NilSize
	} else {
		s += z.Scalar.Msgsize()
	}
	return
}
