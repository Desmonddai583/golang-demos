package common

const (
	// JobSaveDir 任务保存目录
	JobSaveDir = "/cron/jobs/"

	// JobKillerDir 任务强杀目录
	JobKillerDir = "/cron/killer/"

	// JobLockDir 任务锁目录
	JobLockDir = "/cron/lock/"

	// JobWorkerDir 服务注册目录
	JobWorkerDir = "/cron/workers/"

	// JobEventSave 保存任务事件
	JobEventSave = 1

	// JobEventDelete 删除任务事件
	JobEventDelete = 2

	// JobEventKill 强杀任务事件
	JobEventKill = 3
)
