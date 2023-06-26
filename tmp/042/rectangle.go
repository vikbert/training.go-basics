package main

type rectangle struct {
	width  int
	height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) setWidth(w int) {
	(*r).width = w
}
