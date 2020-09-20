package array

import (
	"strings"
)

type Array interface {
	Len() int
}

type IntArray []int
type StrArray []string

func (a IntArray) Len() int { return len(a) }
func (a StrArray) Len() int { return len(a) }

func (a *IntArray) Append(val ...int)    { *a = append(*a, val...) }
func (a *StrArray) Append(val ...string) { *a = append(*a, val...) }

func (a IntArray) IsContain(input int) (contain bool) {
	contain = false
	for _, v := range a {
		if v == input {
			contain = true
			return
		}
	}
	return
}

func (a StrArray) IsContain(input string, ignoreCase bool) (contain bool) {
	contain = false
	for _, v := range a {
		if ignoreCase {
			if strings.EqualFold(input, v) {
				contain = true
				return
			}
		} else {
			if v == input {
				contain = true
				return
			}
		}
	}
	return
}
