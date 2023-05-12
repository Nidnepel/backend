package entity

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
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
	AdminRole     = "admin"
	WorkerRole    = "worker"
	ManagerRole   = "manager"
)
