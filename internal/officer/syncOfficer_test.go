package officer

import (
	"log"
	"testing"
	"time"
)

func TestSyncOfficer(t *testing.T) {
	officerA := NewSyncOfficer("A")
	officerA.AddFunc("A1", func() { log.Println("doing the first task...") })
	officerA.AddFunc("A2", func() { log.Println("doing the second task...") })
	officerA.AddFunc("A3", func() {
		for i := 0; i < 100; i++ {
			go log.Println("doing the third task...", i)
		}
	})
	// officer.Execute()

	officerB := NewSyncOfficer("B")
	officerB.AddFunc("B1", func() { log.Println("B doing the first task...") })
	officerB.AddFunc("B1", func() {
		for i := 0; i < 100; i++ {
			go log.Println("B doing the second task...", i)
		}
	})
	officerB.AddFunc("B1", func() { log.Println("B doing the third task...") })
	officerB.AddFunc("B1", func() { log.Println("B doing the forth task...") })
	// officerB.Execute()

	var officers []Officer
	officers = append(officers, officerA)
	officers = append(officers, officerB)

	for i := 0; i < len(officers); i++ {
		go officers[i].Execute()
	}

	time.Sleep(5 * time.Second)

	t.Error("done")
}
