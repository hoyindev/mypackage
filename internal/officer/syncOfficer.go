package officer

import (
	"log"
	"sync"
)

type SyncOfficer struct {
	ID        string
	Entries   []Entry
	jobWaiter sync.WaitGroup
}

func NewSyncOfficer(id string) *SyncOfficer {
	o := &SyncOfficer{
		ID: id,
	}
	return o
}

func (o *SyncOfficer) AddFunc(jobID string, cmd func()) {
	o.AddJob(jobID, FuncJob(cmd))
}

func (o *SyncOfficer) AddJob(jobID string, cmd Job) {
	entry := &Entry{
		ID:  jobID,
		Job: cmd,
	}
	o.Entries = append(o.Entries, *entry)
}

func (o *SyncOfficer) Execute() {
	log.Println("officer", o.ID, "has", len(o.Entries), "tasks ")
	// o.jobWaiter.Add(len(o.Entries))
	for _, e := range o.Entries {
		o.startJob(e.Job)
	}
	o.jobWaiter.Wait()
}

// startJob runs the given job in a new goroutine.
func (o *SyncOfficer) startJob(j Job) {
	o.jobWaiter.Add(1)
	go func() {
		defer o.jobWaiter.Done()
		j.Run()
	}()
}
