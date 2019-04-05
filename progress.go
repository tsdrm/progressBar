package progressBar

import (
	"sync"
	"time"
	"fmt"
	"strconv"
)

type MODEL_TYPE int

const (
	MODEL_NUMBER  MODEL_TYPE = iota
	MODEL_PROCESS
)

type Progress struct {
	Total   int
	Current int

	isPercentage bool
	Model        MODEL_TYPE
	Prefix       string
	Suffix       string

	Interval time.Duration
	shower   Shower

	c  chan string
	wg *sync.WaitGroup
}

type Shower interface {
	Show(c, t int, prefix, suffix string, isPercentage bool)
	ShowFloatN(c, t, bitSize int, prefix, suffix string, isPercentage bool)
}

func NewBar(total int, model MODEL_TYPE, prefix, suffix string, isPercentage bool) *Progress {
	if total < 1 {
		panic("Total must be greater than 0")
	}

	var progress = &Progress{
		Total:   total,
		Current: 0,

		isPercentage: isPercentage,
		Model:        model,
		Prefix:       prefix,
		Suffix:       suffix,
		shower:       GetShower(model),
		Interval:     100 * time.Millisecond,
		c:            make(chan string, 100),
		wg:           &sync.WaitGroup{},
	}

	return progress
}

func GetShower(model MODEL_TYPE) Shower {
	switch model {
	case MODEL_NUMBER:
		return Number{}
	case MODEL_PROCESS:
		return Bar{}
	default:
		return Bar{}
	}
}

func (p *Progress) Count(n int) int {
	if n < 0 {
		n = 0
	}
	p.Current += n
	p.c <- fmt.Sprintf("%d", p.Current)
	return p.Current
}

func (p *Progress) sleep() {
	time.Sleep(p.Interval)
}

func (p *Progress) Start() {
	go func(progress *Progress) {
		progress.wg.Add(1)
		for {
			progress.shower.Show(progress.Current, progress.Total,
				progress.Prefix, progress.Suffix, progress.isPercentage)
			if progress.Current >= progress.Total {
				break
			}

			var current = <-progress.c
			progress.Current, _ = strconv.Atoi(current)
		}
		time.Sleep(time.Second)
		progress.wg.Done()
	}(p)
}

func (p *Progress) Wait() {
	p.wg.Wait()
}

func (p *Progress) SetInterval(duration time.Duration) {
	p.Interval = duration
}

type ProgressGroup struct {
	Progresses []*Progress
	Interval   time.Duration
}

func (pg *ProgressGroup) sleep() {
	time.Sleep(pg.Interval)
}

func (pg *ProgressGroup) MoveUp(n int) {

}

func (pg *ProgressGroup) MoveDown(n int) {

}

func (pg *ProgressGroup) Start() {

}
