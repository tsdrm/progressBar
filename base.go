package progressBar

import "fmt"

type Base struct {
}

func (b Base) PercentageInt(c, t int) int {
	if t == 0 {
		return 0
	}
	return int(float32(c) / float32(t) * 100)
}

func (b Base) PercentageFloatN(c, t int, n int) string {
	if t == 0 {
		return "0"
	}
	var format = fmt.Sprintf("%df", n)
	return fmt.Sprintf("%."+format, float64(c)/float64(t))
}
