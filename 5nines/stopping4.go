package main

import "fmt"
import "launchpad.net/tomb"

// START OMIT
type Worker struct {
	tomb.Tomb
}

func (w *Worker) run() {
	defer w.Tomb.Done()
	a, b := make(chan bool), make(chan bool)
	close(a); close(b)
	for {
		select {
		case <-a:
			w.Tomb.Kill(fmt.Errorf("a"))
			return
		case <-b:
			w.Tomb.Kill(fmt.Errorf("b"))
			return
		}
	}
}

func main() {
	for i := 0; i < 10; i++ {
		w := &Worker{}
		go w.run()
		fmt.Printf("Worker exited: %v\n", w.Tomb.Wait())
	}
}
// END OMIT
