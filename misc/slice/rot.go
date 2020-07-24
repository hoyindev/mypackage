package slice

//RotLeft  (O(M) time
//e.g. a = 1,2,3,4,5; d = 2
//result = 3,4,5,1,2
func RotLeft(a []int32, d int) []int32 {
	d %= len(a) //d = d % len(a)
	a = append(a, a[:d]...)
	return a[d:]
}
