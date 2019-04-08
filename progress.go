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
	// total progress
	Total int
	// current progress
	Current int

	// whether to use percentage display
	isPercentage bool
	// display style, you can choose percentage or bar
	Model MODEL_TYPE
	// prefix of progress
	Prefix string
	// suffix of progress
	Suffix string

	Interval time.Duration
	shower   Shower

	c  chan string
	wg *sync.WaitGroup
}

type Shower interface {
	// display by integer
	Show(c, t int, prefix, suffix string, isPercentage bool)
	// display by floating point number
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
	if p.Current >= p.Total {
		p.Current = p.Total
	}
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

			var current = <-p.c
			p.Current, _ = strconv.Atoi(current)
		}
		close(p.c)
		fmt.Println()
		progress.wg.Done()
	}(p)
}

func (p *Progress) listen() {
	for {
		var current = <-p.c
		p.Current, _ = strconv.Atoi(current)
		if p.Current == p.Total {
			break
		}
	}
}

func (p *Progress) show() {
	p.shower.Show(p.Current, p.Total, p.Prefix, p.Suffix, p.isPercentage)
}

func (p *Progress) Wait() {
	p.wg.Wait()
}

// Return status of progress, p.Current == p.Total represent process is success.
func (p *Progress) Status() uint8 {
	if p.Current == p.Total {
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

	TotalLine int

	wg *sync.WaitGroup
}

func NewProcessGroup() *ProgressGroup {
	return &ProgressGroup{
		Progresses: []*Progress{},
		Interval:   time.Millisecond * 5,
		TotalLine:  0,

		wg: &sync.WaitGroup{},
	}
}

func (pg *ProgressGroup) Add(p *Progress) {
	pg.Progresses = append(pg.Progresses, p)
	pg.TotalLine ++
}

func (pg *ProgressGroup) sleep() {
	time.Sleep(pg.Interval)
}

func (pg *ProgressGroup) LineMoveUp(n int) {
	fmt.Printf("\033[%dA", n)
}

func (pg *ProgressGroup) LineMoveDown(n int) {
	fmt.Printf("\033[%dB", n)
}

func (pg *ProgressGroup) HideCursor() {
	fmt.Printf("\033[?25l")
}

func (pg *ProgressGroup) ShowCursor() {
	fmt.Printf("\033[?25h")
}

func (pg *ProgressGroup) Start() {
	pg.HideCursor()
	for _, p := range pg.Progresses {
		go func(progress *Progress, progressGroup *ProgressGroup) {
			progressGroup.wg.Add(1)
			defer progressGroup.wg.Done()
			fmt.Println()
			progress.listen()
		}(p, pg)
	}

	var successNum int
	go func() {
		for {
			// check and show
			if successNum == len(pg.Progresses) {
				break
			}
			pg.LineMoveUp(pg.TotalLine)
			successNum = 0
			for _, p := range pg.Progresses {
				if p.Status() == PRO_SUCCESS {
					successNum ++
				}
				p.show()
				fmt.Println()
			}
			pg.sleep()
		}
	}()
}

// Why sleep a second at the end ?
// It is to wait for the display thread to display complete.
// Otherwise it might show something like 99/100.
func (pg *ProgressGroup) Wait() {
	pg.wg.Wait()
	pg.ShowCursor()
	time.Sleep(time.Second)
}
