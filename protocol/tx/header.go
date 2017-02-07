package tx

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

func newHeader(version uint64, results []EntryRef, data EntryRef, minTimeMS, maxTimeMS uint64) *Header {
	h := new(Header)
	h.body.Version = version
	h.body.Results = results
	h.body.Data = data
	h.body.MinTimeMS = minTimeMS
	h.body.MaxTimeMS = maxTimeMS
	return h
}
