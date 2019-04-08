package progressBar

import (
	"testing"
	"time"
)

var sleep = func() {
	time.Sleep(10 * time.Millisecond)
}

var processRun = func(p *Progress, count int, n int, interval time.Duration) {
	for i := 0; i < count; i++ {
		p.Count(n)
		time.Sleep(interval)
	}
}

func TestProgress_Bar(t *testing.T) {
	var count = 123
	var nb *Progress
	{
		nb = NewBar(count, MODEL_NUMBER, "progress: ", "!!!!", false)
		nb.Start()

		for i := 0; i < count; i++ {
			nb.Count(1)
			time.Sleep(10 * time.Millisecond)
		}
		nb.Wait()
	}

	{
		nb = NewBar(count, MODEL_NUMBER, "progress: ", ".", false)
		nb.Start()
		go func() {
			for i := 0; i < count; i++ {
				nb.Count(1)
				time.Sleep(10 * time.Millisecond)
			}
		}()
		time.Sleep(time.Millisecond * 10)
		nb.Wait()
	}

	{
		nb = NewBar(count, MODEL_NUMBER, "progress: ", "!!!!", true)
		nb.Start()

		for i := 0; i < count; i++ {
			nb.Count(1)
			time.Sleep(10 * time.Millisecond)
		}
		nb.Wait()
	}

	{
		nb = NewBar(count, MODEL_NUMBER, "progress: ", ".", true)
		nb.Start()
		go func() {
			for i := 0; i < count; i++ {
				nb.Count(1)
				time.Sleep(10 * time.Millisecond)
			}
		}()
		time.Sleep(time.Millisecond * 10)
		nb.Wait()
	}
}

func TestProcessGroup(t *testing.T) {
	var pg *ProgressGroup
	var count int

	// single process in progress group
	{
		var p1 *Progress
		count = 145

		pg = NewProcessGroup()
		p1 = NewBar(count, MODEL_NUMBER, "A: ", "!!!", false)
		pg.Add(p1)
		pg.Start()
		processRun(p1, count, 2, time.Millisecond*10)
		pg.Wait()

		pg = NewProcessGroup()
		p1 = NewBar(count, MODEL_NUMBER, "B: ", "!!!", false)
		pg.Add(p1)
		pg.Start()
		go processRun(p1, count, 2, time.Millisecond*10)
		sleep()
		pg.Wait()

		pg = NewProcessGroup()
		p1 = NewBar(count, MODEL_NUMBER, "C: ", "!!!", true)
		pg.Add(p1)
		pg.Start()
		processRun(p1, count, 2, time.Millisecond*10)
		pg.Wait()

		pg = NewProcessGroup()
		p1 = NewBar(count, MODEL_NUMBER, "D: ", "!!!", true)
		pg.Add(p1)
		pg.Start()
		go processRun(p1, count, 2, time.Millisecond*10)
		sleep()
		pg.Wait()
	}

	// two process in progress group
	{
		var p1 *Progress
		var p2 *Progress
		count = 123

		pg = NewProcessGroup()
		p1 = NewBar(count, MODEL_NUMBER, "progressA: ", ".", true)
		p2 = NewBar(count, MODEL_NUMBER, "progressB: ", ".", false)
		pg.Add(p1)
		pg.Add(p2)
		pg.Start()
		go processRun(p1, count, 1, time.Millisecond*10)
		go processRun(p2, count, 1, time.Millisecond*10)
		sleep()
		pg.Wait()
	}
}
