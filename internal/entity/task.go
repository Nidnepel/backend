package entity

type Task struct {
	Id             int
	Title          string
	Description    string
	ProgressStatus bool
}

type TaskReportList struct {
	Id       int
	TaskId   int
	ReportId int
}

const (
	ClosedTask = true
	OpenedTask = false
)