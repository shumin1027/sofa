package main

import (
	"fmt"
	"github.com/bitfield/script"
)

func main() {
	//p := script.Exec("cat /Users/shumin/work/xtc/sofa/test.log")
	p := script.Exec("su test01 && bjobs -W")
	var exit int = p.ExitStatus()
	fmt.Println(exit)
	//p.Column(5).Stdout()
	lines, _ := p.CountLines()
	fmt.Println(lines)
	output, _ := p.String()
	fmt.Println(output)
}
