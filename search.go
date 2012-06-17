package main

import (
	"github.com/pmylund/go-bitset"
	"sort"
)

type documents struct {
	tags []bitset.Bitset64
	size uint64
}

type docScore struct {
	doc   uint32
	score uint64
}

type docScores struct {
	store []*docScore
}

func (self docScores) Len() int { return len(self.store) }

func (self docScores) Swap(i, j int) {
	self.store[i], self.store[j] = self.store[j], self.store[i]
}

func (self *docScores) Add(ds *docScore) {
	if self.store == nil {
		self.store = make([]*docScore, 0, 4)
	}
	n := len(self.store)
	if n+1 > cap(self.store) {
		s := make([]*docScore, n, 2*n+1)
		copy(s, self.store)
		self.store = s
	}
	self.store = self.store[0 : n+1]
	self.store[n] = ds
}

type byScore struct{ docScores }

func (self byScore) Less(i, j int) bool {
	return self.docScores.store[i].score < self.docScores.store[j].score
}

func NewDocuments(docs uint32, tags uint64) documents {
	return documents{
		make([]bitset.Bitset64, docs),
		tags}
}

func NewDocument(ids []uint64) *bitset.Bitset64 {
	b := bitset.New64(0)
	for _, id := range ids {
		b.Set(id)
	}
	return b
}

func (self *documents) Set(pos uint32, tag *bitset.Bitset64) {
	self.tags[pos] = *tag
}

func (self *documents) Score(
	master *bitset.Bitset64,
	thresold_ float32) []*docScore {
	thresold := uint64(float64(thresold_) * float64(master.Count()))
	var results docScores = docScores{make([]*docScore, 0, 4)}
	for i, document := range self.tags {
		common := master.Intersection(&document).Count()
		if common > thresold {
			results.Add(&docScore{uint32(i), common})
		}
	}
	sort.Sort(byScore{results})
	return results.store
}
