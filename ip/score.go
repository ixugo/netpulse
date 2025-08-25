package ip

import (
	"expvar"
	"iter"
	"slices"
	"sort"
	"sync"
)

var _ sort.Interface = &StringScorer{}

// StringScorer is used to score successful connection requests. When getting the connection array next time,
// it prioritizes links with higher scores.
type StringScorer struct {
	topScore int
	topStr   string

	data   []string
	scores map[string]int
	length int

	m sync.RWMutex
}

// Len implements sort.Interface.
func (ss *StringScorer) Len() int {
	return ss.length
}

// Less implements sort.Interface.
func (ss *StringScorer) Less(i int, j int) bool {
	return ss.scores[ss.data[i]] > ss.scores[ss.data[j]]
}

// Swap implements sort.Interface.
func (ss *StringScorer) Swap(i int, j int) {
	ss.data[i], ss.data[j] = ss.data[j], ss.data[i]
}

func NewStringScorer(length int) *StringScorer {
	s := StringScorer{
		scores: make(map[string]int, length),
		data:   make([]string, 0, length),
	}
	expvar.Publish("netpulse", expvar.Func(func() interface{} {
		return s.AllWithScores()
	}))

	return &s
}

func (ss *StringScorer) Set(str string) *StringScorer {
	ss.m.Lock()
	defer ss.m.Unlock()

	if _, ok := ss.scores[str]; ok {
		return ss
	}
	ss.scores[str] = 0
	ss.data = append(ss.data, str)
	ss.length = len(ss.data)
	return ss
}

func (ss *StringScorer) Del(str string) {
	ss.m.Lock()
	defer ss.m.Unlock()

	delete(ss.scores, str)
	idx := slices.Index(ss.data, str)
	ss.data = slices.Delete(ss.data, idx, idx+1)
	ss.length = len(ss.data)
}

func (ss *StringScorer) AddScore(str string) {
	ss.m.Lock()
	defer ss.m.Unlock()

	if str == ss.topStr {
		return
	}

	ss.scores[str]++
	if r := ss.scores[str]; r > ss.topScore {
		ss.topScore = r
		ss.topStr = str
		sort.Stable(ss)
	}
}

// All returns an iterator over all strings in the scorer, ordered by score (highest first).
func (ss *StringScorer) All() iter.Seq[string] {
	return func(yield func(string) bool) {
		ss.m.RLock()
		defer ss.m.RUnlock()
		data := ss.data
		for _, str := range data {
			if !yield(str) {
				return
			}
		}
	}
}

// AllWithScores returns an iterator over all string-score pairs in the scorer, ordered by score (highest first).
func (ss *StringScorer) AllWithScores() iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		ss.m.RLock()
		defer ss.m.RUnlock()

		for _, str := range ss.data {
			if !yield(str, ss.scores[str]) {
				return
			}
		}
	}
}
