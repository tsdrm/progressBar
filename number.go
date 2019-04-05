package progressBar

import "fmt"

type Number struct {
	Base
}

func (n Number) Show(c, t int, prefix, suffix string, isPercentage bool) {
	var percentage, total = c, t
	if isPercentage {
		percentage = n.PercentageInt(c, t)
		total = 100
	}
	fmt.Printf("\r%s%d/%d%s", prefix, percentage, total, suffix)
}

func (n Number) ShowFloatN(c, t, bitSize int, prefix, suffix string, isPercentage bool) {
	var percentage = fmt.Sprintf("%."+fmt.Sprintf("%df", bitSize), float32(c)/float32(t))
	fmt.Printf("\r%s%s/%s%s", prefix, percentage, "100", suffix)
}
