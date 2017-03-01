package leetcode

// ByteStack is a stack for byte
type ByteStack []byte

// Push b to ByteStack
func (s *ByteStack) Push(b byte) {
	*s = append(*s, b)
}

// Pop from ByteStack need check Empty() first
func (s *ByteStack) Pop() byte {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

// Empty return true if len(s) == 0
func (s ByteStack) Empty() bool {
	return len(s) == 0
}
