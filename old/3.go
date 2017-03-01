// 3. Isomorphic Strings
// Given two strings s and t, determine if they are isomorphic.
// Two strings are isomorphic if the characters in s can be replaced to get t.
// For example,"egg" and "add" are isomorphic, "foo" and "bar" are not.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsIsomorphic("egg", "add"))
	fmt.Println(IsIsomorphic("foo", "bar"))
}

func IsIsomorphic(s, t string) bool {
	if len(s) != len(t) || len(s) == 0 { // "" is false
		return false
	}

	ref := make(map[byte]byte) // ref[s[i]] = t[i]

	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]
		if v, ok := ref[c1]; ok { // c1 in map
			if v != c2 {
				return false // Conflict
			}
		} else {
			for _, v := range ref {
				if v == c2 {
					return false // c2 already in map
				}
			}
			ref[c1] = c2 // Insert
		}
	}
	return true
}
