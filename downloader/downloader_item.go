package downloader

import (
	"time"
)

type downloadItem interface {
	id() string
	name() string
	createTime() time.Time
	start()
	pause()
	delete()
	state() state
}

type runningState int

const (
	new runningState = iota
	run
	pause
	finished
)

type state struct {
	RunningState  runningState `json:"running_state"`
	TotalBytes    int64        `json:"total_bytes"`
	CompleteBytes int64        `json:"complete_bytes"`
	Rate          float64      `json:"rate"`
	Percent       float64      `json:"percent"`
}
