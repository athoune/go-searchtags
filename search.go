package main

import "github.com/pmylund/go-bitset"

type documents struct {
	tags []bitset.Bitset64
	size uint64
}

func New(docs uint32, tags uint64) documents {
	return documents{
		make([]bitset.Bitset64, docs),
		tags}
}

func (self *documents) Set(pos uint32, tag *bitset.Bitset64) {
	self.tags[pos] = *tag
}

func (self *documents) Score(
	master *bitset.Bitset64, thresold_ float32, iter func(uint32, uint64)) {
	thresold := uint64(float64(thresold_) * float64(self.size))
	for i, document := range self.tags {
		common := master.Intersection(&document).Count()
		if common > thresold {
			iter(uint32(i), common)
		}
	}
}

func main() {

}
