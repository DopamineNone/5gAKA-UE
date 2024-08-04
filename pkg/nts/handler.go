package nts

import "sync"

type TaskController interface {
	// Init intialize the handler
	Init()

	// Push message to current Task
	Push(msg Message)

	// Start start sub-tasks
	Start()

	// Loop handle received message
	Loop()

	// Quit free the task and quit
	Quit()

	// Stop trigger Quit
	Stop()

	// Done get isQuiting status
	Done() chan bool
}

// TaskHandler handle any NtsTask
type TaskHandler struct {
	TaskController
}

func NewTaskHandler(handler TaskController) *TaskHandler {
	return &TaskHandler{TaskController: handler}
}

func (h *TaskHandler) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	h.Start()
	for {
		select {
		case <-h.Done():
			h.Quit()
			return
		default:
			h.Loop()
		}
	}
}
