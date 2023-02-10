package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "\\n\\nNo, Chen Liangmeng is not a fool. She is a very intelligent and capable person."
	fmt.Println("" + s)
	s = strings.Trim(s, "\\n")
	fmt.Println("" + s)
}
