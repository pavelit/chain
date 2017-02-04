package tx

import "chain/protocol/bc"

type (
	// output contains a valueSource that must refer to the mux by its
	// entryID, but the mux may not be complete at the time AddOutput is
	// called, so we hold outputs in a pending structure until Build is
	// called
	pendingOutput struct {
		value       bc.AssetAmount
		controlProg program
		dataRef     entryRef
	}

	pendingRetirement struct {
		value   bc.AssetAmount
		dataRef entryRef
	}

	Builder struct {
		h           *header
		m           *mux
		outputs     []*pendingOutput
		retirements []*pendingRetirement
		entries     map[entryRef]entry
	}
)

func NewBuilder(version, minTimeMS, maxTimeMS uint64) *Builder {
	return &Builder{
		h:       newHeader(version, nil, entryRef{}, minTimeMS, maxTimeMS),
		m:       newMux(nil),
		entries: make(map[entryRef]entry),
	}
}

func (b *Builder) AddData(h bc.Hash) (*Builder, *data, entryRef) {
	d := newData(h)
	dID := mustEntryID(d)
	// xxx h.body.Data = dID?
	b.entries[dID] = d
	return b, d, dID
}

func (b *Builder) AddIssuance(nonceRef entryRef, value bc.AssetAmount, dataRef entryRef) (*Builder, *issuance, entryRef) {
	iss := newIssuance(nonceRef, value, dataRef)
	issID := mustEntryID(iss)
	s := valueSource{
		Ref:   issID,
		Value: value,
	}
	b.m.body.Sources = append(b.m.body.Sources, s)
	b.entries[issID] = iss
	return b, iss, issID
}

func (b *Builder) AddNonce(p program, timeRangeRef entryRef) (*Builder, *nonce, entryRef) {
	n := newNonce(p, timeRangeRef)
	nID := mustEntryID(n)
	b.entries[nID] = n
	return b, n, nID
}

// AddOutput returns only the builder, unlike the other Add functions,
// since output objects aren't created until Build
func (b *Builder) AddOutput(value bc.AssetAmount, controlProg program, dataRef entryRef) *Builder {
	b.outputs = append(b.outputs, &pendingOutput{
		value:       value,
		controlProg: controlProg,
		dataRef:     dataRef,
	})
	return b
}

func (b *Builder) AddRetirement(value bc.AssetAmount, dataRef entryRef) *Builder {
	b.retirements = append(b.retirements, &pendingRetirement{
		value:   value,
		dataRef: dataRef,
	})
	return b
}

func (b *Builder) AddSpend(spentOutput bc.OutputID, value bc.AssetAmount, dataRef entryRef) (*Builder, *spend, entryRef) {
	sp := newSpend(spentOutput, dataRef)
	spID := mustEntryID(sp)
	src := valueSource{
		Ref:   spID,
		Value: value,
	}
	b.m.body.Sources = append(b.m.body.Sources, src)
	b.entries[spID] = sp
	return b, sp, spID
}

func (b *Builder) AddTimeRange(minTimeMS, maxTimeMS uint64) (*Builder, *timeRange, entryRef) {
	tr := newTimeRange(minTimeMS, maxTimeMS)
	trID := mustEntryID(tr)
	b.entries[trID] = tr
	return b, tr, trID
}

func (b *Builder) Build() (entryRef, *header, map[entryRef]entry) {
	muxID := mustEntryID(b.m)
	b.entries[muxID] = b.m
	var n uint64
	for _, po := range b.outputs {
		s := valueSource{
			Ref:      muxID,
			Value:    po.value,
			Position: n,
		}
		n++
		o := newOutput(s, po.controlProg, po.dataRef)
		oID := mustEntryID(o)
		b.entries[oID] = o
		b.h.body.Results = append(b.h.body.Results, oID)
	}
	for _, pr := range b.retirements {
		s := valueSource{
			Ref:      muxID,
			Value:    pr.value,
			Position: n,
		}
		n++
		r := newRetirement(s, pr.dataRef)
		rID := mustEntryID(r)
		b.entries[rID] = r
		b.h.body.Results = append(b.h.body.Results, rID)
	}
	hID := mustEntryID(b.h)
	b.entries[hID] = b.h
	return hID, b.h, b.entries
}

func mustEntryID(e entry) entryRef {
	res, err := entryID(e)
	if err != nil {
		panic(err)
	}
	return res
}
