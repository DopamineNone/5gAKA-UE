package nts

import "sync"

type TaskController interface {
	// Start start sub-tasks
	Start()

	// Loop handle received message
	Loop(msg Message)

	// Quit free the task and quit
	Quit()

	// Stop trigger Quit
	Stop()

	// Done get isQuiting status
	Done() chan bool

	// Take get received message
	Take() chan Message
}

// TaskHandler handle any NtsTask
type TaskHandler struct {
	TaskController
}

func NewTaskHandler(handler TaskController) *TaskHandler {
	return &TaskHandler{TaskController: handler}
}

func (h *TaskHandler) Run(wg *sync.WaitGroup) {
	h.Start()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-h.Done():
				h.Quit()
				return
			case msg := <-h.Take():
				h.Loop(msg)
			}
		}
	}()
}
