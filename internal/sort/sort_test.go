package sort

import (
	"sort"
	"testing"
)

func TestSortByID(t *testing.T) {
	txns := []Txn{
		{"0003", 230},
		{"0012", 20},
		{"0001", 2500},
	}
	sort.Sort(ByID(txns))
	t.Error(txns) //1, 3, 12

	sort.Sort(ByAmt(txns))
	t.Error(txns) //12, 3, 1

	m := map[string]int{"0003": 230, "0012": 20, "0001": 2500}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		t.Error(k, m[k]) //1, 3, 12
	}
}
