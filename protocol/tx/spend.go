package tx

type Spend struct {
	body struct {
		SpentOutput EntryRef
		Data        EntryRef
		ExtHash     extHash
	}
}

func (Spend) Type() string         { return "spend1" }
func (s *Spend) Body() interface{} { return s.body }

func (s *Spend) SpentOutput() EntryRef {
	return s.body.SpentOutput
}

func newSpend(spentOutput, data EntryRef) *Spend {
	s := new(Spend)
	s.body.SpentOutput = spentOutput
	s.body.Data = data
	return s
}
