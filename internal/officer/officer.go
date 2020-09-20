package officer

type Officer interface {
	Execute()
}

type Entry struct {
	ID  string
	Job Job
}

// Job is an interface for submitted cron jobs.
type Job interface {
	Run()
}

// FuncJob is a wrapper that turns a func() into a Job
type FuncJob func()

func (f FuncJob) Run() { f() }
