package shell

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {

	//s := `sh -c 'echo "Hello world"'`
	//cmd := exec.Command("bash", "-c", s)
	//output, err := cmd.CombinedOutput()
	//if err != nil {
	//}
	//fmt.Println(string(output))

	s := "bash -c 'docker ps'"
	//s := `sh -c 'echo "Hello world"'`
	p := Exec(s)
	output, _ := p.String()
	fmt.Println(output)
}
