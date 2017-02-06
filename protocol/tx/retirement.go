package tx

import "chain/protocol/bc"

type Retirement struct {
	body struct {
		Source  valueSource
		DataRef bc.Hash
		ExtHash extHash
	}
}

func (Retirement) Type() string         { return "retirement1" }
func (r *Retirement) Body() interface{} { return r.body }

func newRetirement(source valueSource, dataRef bc.Hash, ordinal int) *Retirement {
	r := new(Retirement)
	r.body.Source = source
	r.body.DataRef = dataRef
	return r
}
