package entity

type Tasklist struct {
	ID       int
	Task     string
	Assignee string
	Deadline string
	Status   bool
}
