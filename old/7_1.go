// 7. Merge Intervals
// Given a collection of intervals, merge all overlapping intervals.

// For example,
// Given [1,3],[2,6],[8,10],[15,18],
// return [1,6],[8,10],[15,18].

package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	start, end int
}

type Intervals []*Interval

func (s Intervals) Len() int { // For sort
	return len(s)
}

func (s Intervals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Intervals) Less(i, j int) bool {
	return s[i].start < s[j].start
}

func (s Intervals) String() (str string) { // For fmt.Println
	for _, i := range s {
		str += fmt.Sprint(*i)
	}
	return
}

func Merge(unMerged Intervals) Intervals {
	if unMerged == nil || len(unMerged) == 0 {
		return nil
	}
	result := make(Intervals, 0, 0)
	sort.Sort(unMerged)

	left, right := unMerged[0].start, unMerged[0].end
	for i := 1; i < len(unMerged); i++ {
		if unMerged[i].start <= right { // Merge
			if unMerged[i].end > right {
				right = unMerged[i].end
			}
		} else {
			result = append(result, &Interval{left, right}) //Append
			left, right = unMerged[i].start, unMerged[i].end
		}
	}
	result = append(result, &Interval{left, right}) // Append the last interval
	return result
}

func main() {
	unMerged := Intervals{&Interval{1, 3}, &Interval{3, 6}, &Interval{8, 10}, &Interval{15, 18}, &Interval{1, 2}}
	merged := Merge(unMerged)
	fmt.Println(merged)
}
