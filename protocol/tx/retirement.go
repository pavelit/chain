package tx

type Retirement struct {
	body struct {
		Source  valueSource
		Data    entryRef
		ExtHash extHash
	}
	ordinal int
}

func (Retirement) Type() string         { return "retirement1" }
func (r *Retirement) Body() interface{} { return r.body }

func (r Retirement) Ordinal() int { return r.ordinal }

func newRetirement(source valueSource, data entryRef, ordinal int) *Retirement {
	r := new(Retirement)
	r.body.Source = source
	r.body.Data = data
	r.ordinal = ordinal
	return r
}
