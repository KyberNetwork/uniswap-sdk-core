package entities

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BaseCurrency) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 6 {
		err = msgp.ArrayError{Wanted: 6, Got: zb0001}
		return
	}
	z.isNative, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "isNative")
		return
	}
	z.isToken, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "isToken")
		return
	}
	z.chainId, err = dc.ReadUint()
	if err != nil {
		err = msgp.WrapError(err, "chainId")
		return
	}
	z.decimals, err = dc.ReadUint()
	if err != nil {
		err = msgp.WrapError(err, "decimals")
		return
	}
	z.symbol, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "symbol")
		return
	}
	z.name, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "name")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BaseCurrency) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 6
	err = en.Append(0x96)
	if err != nil {
		return
	}
	err = en.WriteBool(z.isNative)
	if err != nil {
		err = msgp.WrapError(err, "isNative")
		return
	}
	err = en.WriteBool(z.isToken)
	if err != nil {
		err = msgp.WrapError(err, "isToken")
		return
	}
	err = en.WriteUint(z.chainId)
	if err != nil {
		err = msgp.WrapError(err, "chainId")
		return
	}
	err = en.WriteUint(z.decimals)
	if err != nil {
		err = msgp.WrapError(err, "decimals")
		return
	}
	err = en.WriteString(z.symbol)
	if err != nil {
		err = msgp.WrapError(err, "symbol")
		return
	}
	err = en.WriteString(z.name)
	if err != nil {
		err = msgp.WrapError(err, "name")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BaseCurrency) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 6
	o = append(o, 0x96)
	o = msgp.AppendBool(o, z.isNative)
	o = msgp.AppendBool(o, z.isToken)
	o = msgp.AppendUint(o, z.chainId)
	o = msgp.AppendUint(o, z.decimals)
	o = msgp.AppendString(o, z.symbol)
	o = msgp.AppendString(o, z.name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BaseCurrency) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 6 {
		err = msgp.ArrayError{Wanted: 6, Got: zb0001}
		return
	}
	z.isNative, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "isNative")
		return
	}
	z.isToken, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "isToken")
		return
	}
	z.chainId, bts, err = msgp.ReadUintBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "chainId")
		return
	}
	z.decimals, bts, err = msgp.ReadUintBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "decimals")
		return
	}
	z.symbol, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "symbol")
		return
	}
	z.name, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "name")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BaseCurrency) Msgsize() (s int) {
	s = 1 + msgp.BoolSize + msgp.BoolSize + msgp.UintSize + msgp.UintSize + msgp.StringPrefixSize + len(z.symbol) + msgp.StringPrefixSize + len(z.name)
	return
}
