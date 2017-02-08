package tx

import "chain/protocol/bc"

type Issuance struct {
	body struct {
		AnchorRef bc.Hash
		Value     bc.AssetAmount
		DataRef   bc.Hash
		ExtHash   extHash
	}
}

func (Issuance) Type() string           { return "issuance1" }
func (iss *Issuance) Body() interface{} { return iss.body }

func newIssuance(anchorRef bc.Hash, value bc.AssetAmount, dataRef bc.Hash, ordinal int) *Issuance {
	iss := new(Issuance)
	iss.body.AnchorRef = anchorRef
	iss.body.Value = value
	iss.body.DataRef = dataRef
	return iss
}
