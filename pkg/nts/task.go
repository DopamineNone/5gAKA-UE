package nts

import (
	"context"
	"time"
)

const (
	MinWaitTime = 500
)

// Task base task struct
type Task struct {
	isQuiting chan bool
	msgQueue  chan Message
}

// NewTask return a new NtsTask
func NewTask() *Task {
	return &Task{isQuiting: make(chan bool), msgQueue: make(chan Message)}
}

func (nt *Task) Init() {
	nt.isQuiting = make(chan bool, 1)
	nt.msgQueue = make(chan Message)
}

func (nt *Task) Take(timeout int) Message {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(max(timeout, MinWaitTime))*time.Millisecond)
	defer cancel()
	return nt.poll(ctx)
}

func (nt *Task) Push(msg Message) {
	nt.msgQueue <- msg
}

func (nt *Task) poll(ctx context.Context) Message {
	select {
	case data := <-nt.msgQueue:
		return data
	case <-ctx.Done():
		return Message{
			MessageType: TimeExpired,
			PDU:         nil,
		}
	}
}

func (nt *Task) Start() {
}

func (nt *Task) Loop() {

}

func (nt *Task) Quit() {
	close(nt.msgQueue)
}

func (nt *Task) Stop() {
	close(nt.isQuiting)
}

func (nt *Task) Done() chan bool {
	return nt.isQuiting
}
