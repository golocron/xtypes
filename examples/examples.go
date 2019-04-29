package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/golocron/xtypes"
)

func main() {
	log.SetFlags(0)

	var size int

	flag.IntVar(&size, "size", 10, "Number of tasks in the example")
	flag.Parse()

	if size == 0 {
		log.Printf("you asked to run example with %d tasks. Exiting.", size)
		os.Exit(0)
	}

	rand.Seed(time.Now().UnixNano())

	tasks := make([]*work, size)

	for i := range tasks {
		p := rand.Intn(100 - 1)

		tasks[i] = &work{
			description: fmt.Sprintf("priority: %d", i+p),
			preference:  i + p,
		}
	}

	pqueue := xtypes.NewPriorityQueue(len(tasks))

	for _, t := range tasks {
		if err := pqueue.Push(t); err != nil {
			log.Printf("failed to add work: %s", err)
		}
	}

	log.Printf("doing work in order by priority...")
	if err := doWorkInOrder(pqueue); err != nil {
		log.Printf("failed to do work: %s", err)
	}

	queue := xtypes.NewQueue(len(tasks))

	for _, t := range tasks {
		if err := queue.Push(t); err != nil {
			log.Printf("failed to add work: %s", err)
		}
	}

	log.Printf("doing work in parallel...")
	if err := doWorkInParallel(queue); err != nil {
		log.Printf("failed to do work: %s", err)
	}
}

func doWorkInOrder(queue *xtypes.PriorityQueue) error {
	for !queue.Empty() {
		item, err := queue.Pop()
		if err != nil {
			if err == xtypes.ErrEmptyQueue {
				log.Printf("empty queue")

				return nil
			}

			log.Printf("failed to get item: %s", err)
			continue
		}

		task, ok := item.(*work)
		if !ok {
			log.Printf("failed to get task: invalid type")
			continue
		}

		log.Printf("working in order: %s", task)
	}

	return nil
}

func doWorkInParallel(queue *xtypes.Queue) error {
	items, err := queue.Get(queue.Len())
	if err != nil {
		if err == xtypes.ErrEmptyQueue {
			return nil
		}

		return err
	}

	sema := xtypes.NewSemaphore(runtime.NumCPU())

	var wg sync.WaitGroup

	for _, v := range items {
		task, ok := v.(*work)
		if !ok {
			log.Printf("failed to get task: invalid type")
			continue
		}

		sema.Acquire(1)
		wg.Add(1)

		go func(t *work) {
			defer wg.Done()
			log.Printf("working in parallel: %s", t)
			sema.Release(1)
		}(task)
	}

	wg.Wait()

	return nil
}

type work struct {
	index       int
	preference  int
	description string
}

func (w work) String() string {
	return fmt.Sprintf("task with %s", w.description)
}

func (w *work) Priority() int {
	return w.preference
}

func (w *work) Index() int {
	return w.index
}

func (w *work) SetIndex(idx int) {
	w.index = idx
}
