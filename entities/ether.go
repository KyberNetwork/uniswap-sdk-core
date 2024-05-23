//go:generate go run github.com/tinylib/msgp -unexported -tests=false -v
//msgp:tuple etherInner
//msgp:ignore Ether

package entities

import "github.com/tinylib/msgp/msgp"

// etherInner is the main usage of a 'native' currency, i.e. for Ethereum mainnet and all testnets
type etherInner struct {
	*BaseCurrency
}

// Ether is the main usage of a 'native' currency, i.e. for Ethereum mainnet and all testnets
type Ether struct {
	etherInner
}

func EtherOnChain(chainId uint) *Ether {
	ether := &Ether{etherInner{
		BaseCurrency: &BaseCurrency{
			isNative: true,
			isToken:  false,
			chainId:  chainId,
			decimals: 18,
			symbol:   "ETH",
			name:     "Ether",
		},
	}}
	ether.BaseCurrency.currency = ether
	return ether
}

func (e *Ether) Equal(other Currency) bool {
	v, isEther := other.(*Ether)
	if isEther {
		return v.isNative && v.chainId == e.chainId

	}
	return false
}

func (e *Ether) Wrapped() *Token {
	return WETH9[e.ChainId()]
}

func (e *Ether) DecodeMsg(dc *msgp.Reader) (err error) {
	err = e.etherInner.DecodeMsg(dc)
	if err != nil {
		return
	}
	e.BaseCurrency.currency = e
	return
}

func (e *Ether) UnmarshalMsg(bts []byte) (o []byte, err error) {
	o, err = e.etherInner.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	e.BaseCurrency.currency = e
	return
}
