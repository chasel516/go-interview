package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "gogo123go面试专题goabcgogo"
	p := fmt.Println
	p("1:", strings.TrimPrefix(s, "go")) //go123go面试专题goabcgogo
	p("2:", strings.TrimSuffix(s, "go")) // gogo123go面试专题goabcgo
	p("3:", strings.TrimLeft(s, "go"))   // 123go面试专题goabcgogo
	p("4:", strings.TrimRight(s, "go"))  // gogo123go面试专题goabc
	p("5:", strings.Trim(s, "go"))       // 123go面试专题goabc
	p("6:", strings.TrimFunc(s, func(r rune) bool {
		return r < 128 // trim all ascii chars
	})) // 面试专题
}
