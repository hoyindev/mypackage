package array

import "testing"

func TestArray(t *testing.T) {
	intArr := new(IntArray)
	strArr := new(StrArray)

	intArr.Append(1, 2, 444, -5)
	strArr.Append("a", "abd")

	t.Error(intArr.Len()) //4
	t.Error(strArr.Len()) //2

	t.Error(intArr.IsContain(1))            //true
	t.Error(strArr.IsContain("Abd", true))  //true
	t.Error(strArr.IsContain("Abd", false)) //false
	t.Error(strArr.IsContain("abd", false)) //true

	t.Error("done")
}
