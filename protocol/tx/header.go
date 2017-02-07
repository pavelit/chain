package tx

import "chain/protocol/bc"

type header struct {
	body struct {
		Version              uint64
		ResultRefs           []bc.Hash
		DataRef              bc.Hash
		MinTimeMS, MaxTimeMS uint64
		ExtHash              extHash
	}
}

func (header) Type() string         { return "txheader" }
func (h *header) Body() interface{} { return h.body }

func (header) Ordinal() int { return -1 }

func newHeader(version uint64, resultRefs []bc.Hash, dataRef bc.Hash, minTimeMS, maxTimeMS uint64) *header {
	h := new(header)
	h.body.Version = version
	h.body.ResultRefs = resultRefs
	h.body.DataRef = dataRef
	h.body.MinTimeMS = minTimeMS
	h.body.MaxTimeMS = maxTimeMS
	return h
}
