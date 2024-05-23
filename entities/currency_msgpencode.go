package entities

import (
	"bytes"
	"fmt"

	"github.com/tinylib/msgp/msgp"
)

type currencyDiscriminator byte

const (
	currencyEther currencyDiscriminator = iota
	currencyNative
	currencyToken
)

// EncodeCurrency encode Currency interface as bytes
func EncodeCurrency(currency Currency) []byte {
	if currency == nil {
		return nil
	}

	var (
		encoded = new(bytes.Buffer)
		wr      = msgp.NewWriterBuf(encoded, nil)
	)

	switch currency := currency.(type) {
	case *Ether:
		_ = wr.WriteByte(byte(currencyEther))
		_ = currency.EncodeMsg(wr)
	case *Native:
		_ = wr.WriteByte(byte(currencyNative))
		_ = currency.EncodeMsg(wr)
	case *Token:
		_ = wr.WriteByte(byte(currencyToken))
		_ = currency.EncodeMsg(wr)
	default:
		panic("unsupported Currency concrete type")
	}

	_ = wr.Flush()

	return encoded.Bytes()
}

// DecodeCurrency decodes encoded Currency into Currency interface
func DecodeCurrency(encoded []byte) Currency {
	if encoded == nil {
		return nil
	}

	var (
		rd       = msgp.NewReaderBuf(bytes.NewReader(encoded), nil)
		currency Currency
	)

	discriminator, _ := rd.ReadByte()

	switch currencyDiscriminator(discriminator) {
	case currencyEther:
		ether := new(Ether)
		_ = ether.DecodeMsg(rd)
		currency = ether
	case currencyNative:
		native := new(Native)
		_ = native.DecodeMsg(rd)
		currency = native
	case currencyToken:
		token := new(Token)
		_ = token.DecodeMsg(rd)
		currency = token
	default:
		panic(fmt.Errorf("unsupported Currency discriminator %d", discriminator))
	}

	return currency
}
