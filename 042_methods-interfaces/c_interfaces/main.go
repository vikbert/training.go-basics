package main

import (
	"fmt"
	"local/042_methods-interfaces/c_interfaces/interfaces"
)

func main() {
	fmt.Println("Interfaces:")

	fmt.Println("\nFirst case:", "\n--------")
	var f1 interfaces.Formatter = BinaryFormatter{}
	var p1 interfaces.Parser = BinaryParser{}
	check(42, f1, p1)

	fmt.Println("\nSecond case:", "\n--------")
	var f2 interfaces.Formatter = BarFormatter{}
	var p2 interfaces.Parser = BarParser{}
	check(42, f2, p2)

	fmt.Println("\nLoggerParser case:", "\n--------")
	var f3 interfaces.Formatter = BinaryFormatter{}
	var p3 interfaces.Parser = LoggerParser{}
	check(42, f3, p3.(interfaces.Parser))
}

func check(n int, f interfaces.Formatter, p interfaces.Parser) {
	out := p.Parse(f.Format(n))
	if out != n {
		panic(fmt.Sprintf("got %d, wanted %d", out, n))
	}
}
