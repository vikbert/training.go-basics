package main

var (
	prev     = 0
	prevPrev = 0
)

func Reset() {
	prev, prevPrev = 0, 0
}

func NextFib() int {
	if prev == 0 && prevPrev == 0 {
		prev = 1
		return 1
	}

	res := prev + prevPrev
	prev, prevPrev = res, prev
	return res
}
