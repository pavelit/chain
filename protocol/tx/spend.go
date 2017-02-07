package tx

import "chain/protocol/bc"

type Spend struct {
	body struct {
		SpentOutput bc.OutputID
		Data        EntryRef
		ExtHash     extHash
	}
}

func (Spend) Type() string         { return "spend1" }
func (s *Spend) Body() interface{} { return s.body }

func newSpend(spentOutput bc.OutputID, data EntryRef) *Spend {
	s := new(Spend)
	s.body.SpentOutput = spentOutput
	s.body.Data = data
	return s
}
