package entity

type Project struct {
	Id     int
	Title  string
	Status bool
}

const (
	OpenedProject = true
	ClosedProject = false
)
