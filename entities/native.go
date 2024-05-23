//go:generate go run github.com/tinylib/msgp -unexported -tests=false -v
//msgp:tuple nativeInner
//msgp:ignore Native

package entities

import "github.com/tinylib/msgp/msgp"

type nativeInner struct {
	*BaseCurrency
	wrapped *Token
}

type Native struct {
	nativeInner
}

func NewNative(wrapped *Token, symbol, name string) Currency {
	native := &Native{nativeInner{
		BaseCurrency: &BaseCurrency{
			isNative: true,
			isToken:  false,
			chainId:  wrapped.ChainId(),
			decimals: wrapped.Decimals(),
			symbol:   symbol,
			name:     name,
		},
		wrapped: wrapped,
	}}
	native.BaseCurrency.currency = native
	return native
}

func (n *Native) Equal(other Currency) bool {
	v, isNative := other.(*Native)
	if isNative {
		return v.isNative && v.chainId == n.chainId

	}
	return false
}

func (n *Native) Wrapped() *Token {
	return n.wrapped
}

func (n *Native) DecodeMsg(dc *msgp.Reader) (err error) {
	err = n.nativeInner.DecodeMsg(dc)
	if err != nil {
		return
	}
	n.BaseCurrency.currency = n
	return
}

func (n *Native) UnmarshalMsg(bts []byte) (o []byte, err error) {
	o, err = n.nativeInner.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	n.BaseCurrency.currency = n
	return
}
