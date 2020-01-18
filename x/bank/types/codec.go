package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// generic codec to be used in this module
var ModuleCdc *codec.Codec

func RegisterCodec(cdc *codec.Codec) {

}

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
