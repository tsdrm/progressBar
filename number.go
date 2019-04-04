package progressBar

type Number struct {
	Base
}

func (n Number) Show(c, t int, prefix, suffix string) {
	n.PercentageInt(c, t)
}
