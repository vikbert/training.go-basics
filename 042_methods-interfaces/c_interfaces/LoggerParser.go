package main

import (
	"fmt"
	"local/042_methods-interfaces/c_interfaces/interfaces"
)

type LoggerParser struct {
	interfaces.Parser
}

func (lp LoggerParser) Parse(target string) int {
	p := BinaryParser{}
	fmt.Printf(">>>>>\n[INFO] string to be parsed: %s\n<<<<<\n", target)

	return p.Parse(target)
}
