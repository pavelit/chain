package tx

import "chain/protocol/bc"

type Issuance struct {
	body struct {
		Anchor  entryRef
		Value   bc.AssetAmount
		Data    entryRef
		ExtHash extHash
	}
	ordinal int
}

func (Issuance) Type() string           { return "issuance1" }
func (iss *Issuance) Body() interface{} { return iss.body }

func (iss Issuance) Ordinal() int { return iss.ordinal }

func newIssuance(anchor entryRef, value bc.AssetAmount, data entryRef, ordinal int) *Issuance {
	iss := new(Issuance)
	iss.body.Anchor = anchor
	iss.body.Value = value
	iss.body.Data = data
	iss.ordinal = ordinal
	return iss
}
