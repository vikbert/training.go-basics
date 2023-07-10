package main

import (
	"fmt"
	"local/042_methods-interfaces/c_interfaces/interfaces"
)

type LoggerParser struct {
	interfaces.Parser
}

func (lp LoggerParser) Parse(target string) int {
	fmt.Printf(">>>>>\n[INFO] string to be parsed: %s\n<<<<<\n", target)

	return lp.Parser.Parse(target)
}
