package tx

import "chain/protocol/bc"

type nonce struct {
	body struct {
		Program      program
		TimeRangeRef bc.Hash
		ExtHash      extHash
	}
}

func (nonce) Type() string         { return "nonce1" }
func (n *nonce) Body() interface{} { return n.body }

func (nonce) Ordinal() int { return -1 }

func newNonce(p program, trRef bc.Hash) *nonce {
	n := new(nonce)
	n.body.Program = p
	n.body.TimeRangeRef = trRef
	return n
}
