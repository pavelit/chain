package tx

import "chain/protocol/bc"

type Issuance struct {
	body struct {
		Anchor  EntryRef
		Value   bc.AssetAmount
		Data    EntryRef
		ExtHash extHash
	}
}

func (Issuance) Type() string           { return "issuance1" }
func (iss *Issuance) Body() interface{} { return iss.body }

func (iss *Issuance) AssetID() bc.AssetID {
	return iss.body.Value.AssetID
}

func (iss *Issuance) Amount() uint64 {
	return iss.body.Value.Amount
}

func newIssuance(anchor EntryRef, value bc.AssetAmount, data EntryRef) *Issuance {
	iss := new(Issuance)
	iss.body.Anchor = anchor
	iss.body.Value = value
	iss.body.Data = data
	return iss
}
