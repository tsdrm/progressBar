package progressBar

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
	Show(c, t int)
}

func NewBar(total int, model MODEL_TYPE, prefix, suffix string) *Progress {
	if total < 1 {
		panic("Total must be greater than 0")
	}
	return &Progress{
		Total:   total,
		Current: 0,

		Model:  model,
		Prefix: prefix,
		Suffix: suffix,
	}
}

func (b *Progress) Percentage() int {
	if b.Current == 0 {
		return 0
	}
	return b.Current / b.Total
}

func (b *Progress) Count(n int) int {
	if n < 0 {
		n = 0
	}
	b.Current += n
	return b.Current
}

func (b *Progress) Start() {
	go func() {
		c, t := b.Current, b.Total
		if b.isPercentage {
			c, t = b.Percentage(), 100
		}
		b.shower.Show(c, t)
	}()
}
