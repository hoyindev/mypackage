package officer

import (
	"log"
	"testing"
)

func TestOfficer(t *testing.T) {
	officer := NewAsyncOfficer("A")
	officer.AddFunc("A1", func() { log.Println("doing the first task...") })
	officer.AddFunc("A2", func() { log.Println("doing the second task...") })
	officer.AddFunc("A3", func() { log.Println("doing the third task...") })
	officer.Execute()

	officerB := NewAsyncOfficer("B")
	officerB.AddFunc("B1", func() { log.Println("doing the first task...") })
	officerB.Execute()
	t.Error("done")
}
