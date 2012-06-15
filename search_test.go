package main

import (
	"fmt"
	"github.com/pmylund/go-bitset"
	"math/rand"
	"testing"
)

func TestSearch(t *testing.T) {
	var DOCS_SIZE uint32 = 50
	var TAGS_SIZE uint64 = 10
	docs := New(DOCS_SIZE, TAGS_SIZE)
	var i uint32
	for i = 0; i < DOCS_SIZE; i++ {
		d := bitset.New64(TAGS_SIZE)
		d.Set(10)
		d.Set(uint64(rand.Int63n(int64(TAGS_SIZE))))
		docs.Set(i, d)
	}
	fmt.Println(docs)
	d := bitset.New64(TAGS_SIZE)
	d.Set(10)
	d.Set(uint64(rand.Int63n(int64(TAGS_SIZE))))
	docs.Score(d, 0, func(id uint32, score uint64) {
		fmt.Println(id, score)
	})
}
