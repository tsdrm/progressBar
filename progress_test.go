package progressBar

import (
	"testing"
	"time"
	"sync"
	"fmt"
)

func TestProgress_Bar(t *testing.T) {
	var nb = NewBar(100, MODEL_NUMBER, "progress: ", "!!!!", false)
	nb.Start()

	for i := 0; i < 100; i++ {
		nb.Count(1)
		time.Sleep(10 * time.Millisecond)
	}
	nb.Wait()
	return
}

func TestProcessGroup(t *testing.T) {
	var pg = NewProcessGroup()
	pg.Add(NewBar(100, MODEL_NUMBER, "A: ", "!!!", false))
	pg.Add(NewBar(100, MODEL_NUMBER, "B: ", "!!!", false))
	pg.Start()

	pg.Wait()

	wg := &sync.WaitGroup{}
	go func(l *sync.WaitGroup) {
		l.Add(1)
		defer func() {
			l.Done()
			fmt.Println("-=--=-=-=--===-=")
		}()
		fmt.Println(11111)
		time.Sleep(time.Second * 3)
		fmt.Println(22222)
	}(wg)
	wg.Add(1)
	wg.Wait()
}
