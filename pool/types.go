package pool

import (
	"context"
	"time"
)

type TaskPool interface {
	//Start 开始调度任务
	Start() error
	// shutdown 关闭任务池
	ShutDown() (<-chan struct{}, error)
	//ShutDownNow 立即关闭任务池
	ShutDownNow() ([]Task, error)
	// States 暴露 TaskPool 生命周期内的运行状态
	// ctx 是让用户来控制什么时候退出采样。那么最基本的两个退出机制：一个是 ctx 被 cancel 了或者超时了，一个是TaskPool 被关闭了
	// error 仅仅表示创建 chan state 是否成功
	// interval 表示获取TaskPool运行期间内部状态的周期/时间间隔
	States(ctx context.Context, interval time.Duration) (<-chan State, error)
	Submit(ctx context.Context, task Task) error
}

// Task 代表一个任务
type Task interface {
	// Run 执行任务
	// 如果 ctx 设置了超时时间，那么实现者需要自己决定是否进行超时控制
	Run(ctx context.Context) error
}

type State struct {
	PoolState       int32
	GoCnt           int32
	WaitingTasksCnt int
	QueueSize       int
	RunningTasksCnt int32
	Timestamp       int64
}
