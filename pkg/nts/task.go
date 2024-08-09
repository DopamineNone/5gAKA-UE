package nts

const (
	MinWaitTime = 500
)

// Task base task struct
type Task struct {
	isQuiting  chan bool
	isPrepared chan bool
	msgQueue   chan Message
}

// NewTask return a new NtsTask
func NewTask() *Task {
	return &Task{isQuiting: make(chan bool, 1), isPrepared: make(chan bool, 1), msgQueue: make(chan Message, 1)}
}

func (nt *Task) Start() {
	close(nt.isPrepared)
}

func (nt *Task) Loop(msg Message) {

}

func (nt *Task) Quit() {
	close(nt.isPrepared)
	close(nt.msgQueue)
}

func (nt *Task) Stop() {
	close(nt.isQuiting)
}

func (nt *Task) PushMessage(msg Message) {
	<-nt.isPrepared
	select {
	case <-nt.isQuiting:
		return
	default:
		nt.msgQueue <- msg
	}
}

func (nt *Task) Done() chan bool {
	return nt.isQuiting
}

func (nt *Task) Take() chan Message {
	return nt.msgQueue
}
