package tx

import "chain/protocol/bc"

type valueSource struct {
	Ref      entryRef
	Value    bc.AssetAmount
	Position uint64 // zero unless Ref is a mux
}

type valueDestination struct {
	Ref      entryRef
	Position uint64
}
