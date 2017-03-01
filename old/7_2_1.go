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

func Insert(in *Interval, list *Intervals) {
	old := *list
	i, j, n := 0, 0, len(old)
	left, right := in.start, in.end
	for i = 0; i < n; i++ { // Find i
		if in.start < old[i].start { // start[]
			break
		} else if in.start <= old[i].end { // [start]
			left = old[i].start
			break
		}
	}
	// If in.start > old[n-1].end, the breaks won't happen.
	// Then, i=n, j=n, the next loop won't work.
	for j = i; j < n; j++ { //Find j
		if in.end < old[j].start { // end []
			j -= 1
			break
		} else if in.end <= old[j].end { // [end]
			right = old[j].end
			break
		}
	}
	// If in.end > old[n-1].end, the last j++ has been done, j=n
	if j == n {
		j -= 1
	}
	// *list = append((*list)[:i], (*list)[j+1:]...) // Delete: save elements before i and behind j
	prefix, suffix := old[:i], old[j+1:]
	temp := append(Intervals{&Interval{left, right}}, suffix...)
	// *list = append(prefix, &Interval{left, right})
	// *list = append(*list, suffix...)
	// *list = append(*list, &Interval{left, right})
	*list = append(prefix, temp...)
}

func main() {
	list := &Intervals{&Interval{3, 5}, &Interval{6, 7}, &Interval{10, 15}}
	in := &Interval{3, 19}
	Insert(in, list)
	fmt.Println(list)
}
