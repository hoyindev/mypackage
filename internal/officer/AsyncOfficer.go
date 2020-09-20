package officer

import (
	"log"
)

type AsyncOfficer struct {
	ID      string
	Entries []Entry
}

func NewAsyncOfficer(id string) *AsyncOfficer {
	o := &AsyncOfficer{
		ID: id,
	}
	return o
}

func (o *AsyncOfficer) AddFunc(jobID string, cmd func()) {
	o.AddJob(jobID, FuncJob(cmd))
}

func (o *AsyncOfficer) AddJob(jobID string, cmd Job) {
	entry := &Entry{
		ID:  jobID,
		Job: cmd,
	}
	o.Entries = append(o.Entries, *entry)
}

func (o *AsyncOfficer) Execute() {
	log.Println("officer", o.ID, "has", len(o.Entries), "tasks ")
	for _, e := range o.Entries {
		e.Job.Run()
	}
}
