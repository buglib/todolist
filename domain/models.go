package domain

// 定义任务状态
const (
	Todo = iota
	Done
)

type Task struct {
	ID       uint   `json:"id"`
	TaskInfo string `json:"taskInfo"`
	State    uint   `json:"state"`
}
