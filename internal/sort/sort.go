package sort

type Txn struct {
	ID  string
	Amt int
}

// BuID implements sort.Interface based on the ID field.
type ByID []Txn
type ByAmt []Txn

func (a ByID) Len() int           { return len(a) }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a ByAmt) Len() int           { return len(a) }
func (a ByAmt) Less(i, j int) bool { return a[i].Amt < a[j].Amt }
func (a ByAmt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
