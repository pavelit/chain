package tx

import "chain/protocol/bc"

type Output struct {
	body struct {
		Source         valueSource
		ControlProgram bc.Program
		DataRef        bc.Hash
		ExtHash        extHash
	}
}

func (Output) Type() string         { return "output1" }
func (o *Output) Body() interface{} { return o.body }

func newOutput(source valueSource, controlProgram bc.Program, dataRef bc.Hash, ordinal int) *Output {
	out := new(Output)
	out.body.Source = source
	out.body.ControlProgram = controlProgram
	out.body.DataRef = dataRef
	return out
}
