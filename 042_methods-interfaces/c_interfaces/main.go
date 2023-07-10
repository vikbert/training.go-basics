package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Parser interface {
	parse(string) int
}

type Formatter interface {
	format(int) string
}

type BinaryFormatter struct{}
type BinaryParser struct{}

func (bf BinaryParser) parse(s string) int {
	value, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic("Faided to convert string to binary")
	}

	fmt.Println("binary value", value)
	return int(value)
}

func (bf BinaryFormatter) format(i int) string {
	formatInt := strconv.FormatInt(int64(i), 2)
	fmt.Println("int value", formatInt)
	return formatInt
}

func main() {
	fmt.Println("Interfaces:")
	var formatter Formatter = BinaryFormatter{}
	var parser Parser = BinaryParser{}

	check(8, formatter, parser)

	printMethods(parser)
	fmt.Println("Done")
}

func check(n int, f Formatter, p Parser) {
	out := p.parse(f.format(n))
	if out != n {
		panic(fmt.Sprintf("got %d, wanted %d", out, n))
	}
}

func printMethods(target any) {
	t := reflect.TypeOf(&target)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
	}
}
