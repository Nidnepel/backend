package entity

type Project struct {
	Id          int
	Title       string
	Description string
	Status      bool
}

const (
	OpenedProject = true
	ClosedProject = false
)
