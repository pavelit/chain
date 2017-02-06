package tx

import "chain/protocol/bc"

type Spend struct {
	body struct {
		SpentOutput bc.OutputID
		DataRef     bc.Hash
		ExtHash     extHash
	}
}

func (Spend) Type() string         { return "spend1" }
func (s *Spend) Body() interface{} { return s.body }

func newSpend(spentOutput bc.OutputID, dataRef bc.Hash, ordinal int) *Spend {
	s := new(Spend)
	s.body.SpentOutput = spentOutput
	s.body.DataRef = dataRef
	return s
}
