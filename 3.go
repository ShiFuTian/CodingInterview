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
		if c1, ok := ref[s[i]]; !ok { // K not in map
			for _, c2 := range ref { // V in map?
				if c2 == t[i] {
					return false
				}
			}
			ref[s[i]] = t[i] // Add
		} else if c1 != t[i] { // K in map, but conflict
			return false
		}
	}

	return true
}
