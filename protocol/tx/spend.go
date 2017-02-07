package tx

type Spend struct {
	body struct {
		SpentOutput entryRef // must be an Output entry
		Data        entryRef // must be a Data entry
		ExtHash     extHash
	}
	ordinal int
}

func (Spend) Type() string         { return "spend1" }
func (s *Spend) Body() interface{} { return s.body }

func (s Spend) Ordinal() int { return s.ordinal }

func newSpend(spentOutput, data entryRef, ordinal int) Entry {
	s := new(Spend)
	s.body.SpentOutput = spentOutput
	s.body.Data = data
	s.ordinal = ordinal
	return s
}
