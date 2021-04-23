package ch6_method

import (
	"bytes"
	"fmt"
	"math"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, x%64
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, x%64
	return len(s.words)>word && s.words[word]&(1<<bit)!=0
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i := range t.words {
		if i < len(s.words) {
			s.words[i] |= t.words[i]
		} else {
			s.words = append(s.words, t.words[i])
		}
	}
}

func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, w := range s.words {
		if w == 0 {
			continue
		}
		for offset:=0; offset<64; offset++ {
			if w & (1<<offset) != 0 {
				if buf.Len() > len("{") {
					buf.WriteString(", ")
				}
				_, _ = fmt.Fprintf(&buf, "%d", 64*i+offset)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	l := 0
	for _, w := range s.words {
		if w == 0 { continue }
		for off:=0; off<64; off++ {
			if w&(1<<off) != 0 {
				l++
			}
		}
	}
	return l
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	if word >= len(s.words) { return }
	s.words[word] &= math.MaxUint64-(1<<bit)
}

func (s *IntSet) Clear(x int) {
	s.words = make([]uint64, 0)
}

func (s *IntSet) Copy() *IntSet {
	ret := &IntSet{make([]uint64, len(s.words))}
	for i := range s.words {
		ret.words[i] = s.words[i]
	}
	return ret
}

func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		s.words[i] ^= t.words[i]
		s.words[i] &= s.words[i]
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	if len(s.words) < len(t.words) {
		s.words, t.words = t.words, s.words
	}

	for i := range t.words {
		s.words[i] ^= t.words[i]
	}
}

func (s *IntSet) Elems() []int {
	ret := make([]int, 0, s.Len())
	for i, w := range s.words {
		if w == 0 {
			continue
		}
		for off:=0; off<64; off++ {
			if w&(1<<off) != 0 {
				ret = append(ret, i*64+off)
			}
		}
	}
	return ret
}
