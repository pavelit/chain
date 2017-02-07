package tx

type Output struct {
	body struct {
		Source         valueSource
		ControlProgram bc.Program
		Data           EntryRef
		ExtHash        extHash
	}
}

func (Output) Type() string         { return "output1" }
func (o *Output) Body() interface{} { return o.body }

func newOutput(source valueSource, controlProgram bc.Program, data EntryRef) *Output {
	out := new(Output)
	out.body.Source = source
	out.body.ControlProgram = controlProgram
	out.body.Data = data
	return out
}
