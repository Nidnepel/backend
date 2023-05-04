package entity

type User struct {
	Id       int
	Login    string
	Password string
	Role     string
	Status   bool
}

type UserProjectList struct {
	Id        int
	UserId    int
	ProjectId int
}

type UserTaskList struct {
	Id        int
	UserId    int
	TaskId    int
	ProjectId int
}

const (
	ActiveUser    = true
	NotActiveUser = false
	AdminRole     = "Admin"
	WorkerRole    = "Worker"
	ManagerRole   = "Manager"
)
