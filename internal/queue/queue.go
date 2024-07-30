package queue

import (
	"log/slog"
	"time"
)

type Worker struct {
	id      uint8
	slogger slog.Logger
}

func (w Worker) ProcessMessage(p Payload) {
	w.slogger.Info("Processing message", "message_id", p.id, "worker_id", w.id)
	time.Sleep(2 * time.Second)
	w.slogger.Info("Done processing message", "message_id", p.id, "worker_id", w.id)
}

type Payload struct {
	id      uint8
	message string
}

func queues() {
	// var workers []Worker = []Worker{{1}, {2}, {3}}

	// var messageQueue = make(chan Payload, 10)

	// messageQueue <- Payload{1, "Hello!"}
	// messageQueue <- Payload{2, "Bye!"}
	// messageQueue <- Payload{3, "Ciao!"}

	// var wg sync.WaitGroup

	// for _, worker := range workers {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		worker.ProcessMessage(<-messageQueue)
	// 	}()
	// }

	// wg.Wait()
}
