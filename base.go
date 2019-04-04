package progressBar

import "fmt"

type Base struct {
}

func (b Base) PercentageInt(c, t int) int {
	if c == 0 {
		return 0
	}
	return c / t
}

func (b Base) PercentageFloatN(c, t int, n int) string {
	if c == 0 {
		return "0"
	}
	var format = fmt.Sprintf("%df", n)
	return fmt.Sprintf("%."+format, float64(c)/float64(t))
}
