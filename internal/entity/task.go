package entity

type Task struct {
	Id             int
	Title          string
	Description    string
	ProgressStatus bool `db:"progress_status"`
}

type TaskReportList struct {
	Id       int
	TaskId   int
	ReportId int
}

const (
	ClosedTask = false
	OpenedTask = true
)
