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

const (
	PRO_SUCCESS uint8 = iota
	PRO_RUNNING
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
		fmt.Println("")
		progress.wg.Done()
	}(p)
}

func (p *Progress) Wait() {
	p.wg.Wait()
}

func (p *Progress) Status() uint8 {
	if p.Current >= p.Total {
		return PRO_SUCCESS
	}
	return PRO_RUNNING
}

func (p *Progress) SetInterval(duration time.Duration) {
	p.Interval = duration
}

type ProgressGroup struct {
	Progresses []*Progress
	Interval   time.Duration

	CurrentLine int
	TotalLine   int
	SuccessNum  int

	wg *sync.WaitGroup
}

func NewProcessGroup() *ProgressGroup {
	return &ProgressGroup{
		Progresses:  []*Progress{},
		Interval:    time.Millisecond * 5,
		CurrentLine: 0,
		TotalLine:   0,

		SuccessNum: 0,
		wg:         &sync.WaitGroup{},
	}
}

func (pg *ProgressGroup) Add(p *Progress) {
	pg.Progresses = append(pg.Progresses, p)
}

func (pg *ProgressGroup) sleep() {
	time.Sleep(pg.Interval)
}

func (pg *ProgressGroup) MoveUp(n int) {

}

func (pg *ProgressGroup) MoveDown(n int) {

}

func (pg *ProgressGroup) Start() {
	for _, p := range pg.Progresses {
		go func(progress *Progress, progressGroup *ProgressGroup) {
			progressGroup.wg.Add(1)
			defer progressGroup.wg.Done()

			fmt.Println("--------------")
			progress.Start()
			fmt.Println("==============")
			progress.Wait()

		}(p, pg)
	}

	//go func() {
	//	for {
	//
	//		pg.sleep()
	//	}
	//}()
}

func (pg *ProgressGroup) Wait() {
	pg.wg.Wait()
}
