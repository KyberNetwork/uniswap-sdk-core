package entities

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	assert.True(t, EtherOnChain(1).Equal(EtherOnChain(1)), "ether on same chains is ether")
	assert.False(t, EtherOnChain(1).Equal(t0), "ether is not token0")
	assert.False(t, t1.Equal(t0), "token1 is not token0")
	assert.True(t, t0.Equal(t0), "token0 is token0")
	assert.True(t, t0.Equal(NewToken(1, AddressZero, 18, "symbol", "name")), "token0 is equal to another token0")
}

func TestMsgpEndecode(t *testing.T) {
	for _, currency := range []Currency{
		EtherOnChain(1),
		NewNative(NewToken(1, common.Address{}, 18, "WCNATIVE", "Wrapped Custom Native Token"), "CNATIVE", "Custom Native Token"),
		NewToken(3, AddressOne, 18, "", ""),
	} {
		encoded := EncodeCurrency(currency)
		decoded := DecodeCurrency(encoded)
		assert.EqualValues(t, currency, decoded)
	}
}
