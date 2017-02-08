package tx

import "chain/protocol/bc"

type Header struct {
	body struct {
		Version              uint64
		Results              []EntryRef
		Data                 EntryRef
		MinTimeMS, MaxTimeMS uint64
		ExtHash              extHash
	}
}

func (Header) Type() string         { return "txheader" }
func (h *Header) Body() interface{} { return h.body }

// Inputs returns all input entries (as two lists: spends and
// issuances) reachable from a header's result entries.
func (h *Header) Inputs() (spends, issuances []EntryRef, err error) {
	sMap := make(map[bc.Hash]EntryRef)
	iMap := make(map[bc.Hash]EntryRef)

	var accum func(EntryRef)
	accum = func(ref EntryRef) {
		switch e := ref.Entry.(type) {
		case *Spend:
			hash, err := ref.Hash()
			if err != nil {
				return nil, nil, err
			}
			sMap[hash] = ref

		case *Issuance:
			hash, err := ref.Hash()
			if err != nil {
				return nil, nil, err
			}
			iMap[hash] = ref

		case *mux:
			for _, s := range e.body.Sources {
				accum(s.Ref)
			}
		}
	}

	for _, r := range h.body.Results {
		accum(r)
	}

	for _, e := range sMap {
		spends = append(spends, e)
	}
	for _, e := range iMap {
		issuances = append(issuances, e)
	}
	return spends, issuances, nil
}

func newHeader(version uint64, results []EntryRef, data EntryRef, minTimeMS, maxTimeMS uint64) *Header {
	h := new(Header)
	h.body.Version = version
	h.body.Results = results
	h.body.Data = data
	h.body.MinTimeMS = minTimeMS
	h.body.MaxTimeMS = maxTimeMS
	return h
}
