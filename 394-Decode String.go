package leetcode

import "bytes"

// Furthermore, you may assume that the original data does not contain any digits
// and that digits are only for those repeat numbers, k.
// For example, there won't be input like 3a or 2[4].

func decodeString(s string) string {
	raw := []byte(s)
	ans, _ := decode(raw, 0, len(raw))
	return ans.String()
}

func decode(raw []byte, left, right int) (*bytes.Buffer, int) {
	var buf bytes.Buffer

	var k int

	for i := left; i < right; i++ {
		c := raw[i]

		switch {
		case c >= 48 && c <= 57: // 0-9
			k = k*10 + int(c-48)

		case c >= 97 && c <= 122: // a-z
			buf.WriteByte(c)

		case c == 91: // [
			sub, end := decode(raw, i+1, right) // start at next byte
			for j := 0; j < k; j++ {
				buf.Write(sub.Bytes())
			}
			i = end // deliver end to i
			k = 0   // resert k

		case c == 93: // ]
			return &buf, i // end decode at index of ']'
		}
	}

	return &buf, right // end at the last byte
}
