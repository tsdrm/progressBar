package progressBar

import (
	"sync"
	"time"
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

	shower Shower
}

type Shower interface {
	Show(c, t int, prefix, suffix string)
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

func (b *Progress) Count(n int) int {
	if n < 0 {
		n = 0
	}
	b.Current += n
	return b.Current
}

func (b *Progress) Start() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		b.shower.Show(b.Current, b.Total, b.Prefix, b.Suffix)
	}()
	wg.Wait()
}

type ProgressGroup struct {
	Progresses []*Progress
	Interval   time.Duration
}

func (pg *ProgressGroup) sleep() {
	time.Sleep(pg.Interval * time.Millisecond)
}

func (pg *ProgressGroup) MoveUp(n int) {

}

func (pg *ProgressGroup) MoveDown(n int) {

}

func (pg *ProgressGroup) Start() {

}
