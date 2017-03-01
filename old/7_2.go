// 7. Insert Interval
// Given a set of non-overlapping & sorted intervals,
// insert a new interval into the intervals (merge if necessary).

// Example 1:
// Given intervals [1,3],[6,9], insert and merge [2,5] in as [1,5],[6,9].

// Example 2:
// Given [1,2],[3,5],[6,7],[8,10],[12,16], insert and merge [4,9] in as [1,2],[3,10],[12,16].

// This is because the new interval [4,9] overlaps with [3,5],[6,7],[8,10].

package main

import (
	"fmt"
)

type Interval struct {
	start, end int
}

type Intervals []*Interval

func (s Intervals) String() (str string) { // For fmt.Println
	for _, i := range s {
		str += fmt.Sprint(*i)
	}
	return
}

func Insert(in *Intervals, it *Interval) (out *Intervals) {
	out = new(Intervals)
	for _, this := range *in {
		if it.start > this.end { // No overlap with it
			*out = append(*out, this)
		} else if it.end < this.start { // No overlap with it
			*out = append(*out, this)
		} else { // Overlap
			it = &Interval{min(this.start, it.start), max(this.end, it.end)}
		}
	}
	*out = append(*out, it)
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := &Intervals{&Interval{3, 5}, &Interval{6, 7}, &Interval{10, 15}}
	it := &Interval{1, 2}
	out := Insert(in, it)
	fmt.Println(out)
}
