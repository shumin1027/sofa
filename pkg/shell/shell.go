package shell

import (
	"bytes"
	"github.com/bitfield/script"
	"os/exec"
)

// 对 github.com/bitfield/script 的包装
// 解决无法执行类似 "bash -c 'docker ps'" 这样的命令的问题
type PipeWrap struct {
	*script.Pipe
}

func (p *PipeWrap) ToPipe(s string) *script.Pipe {
	return p.Pipe
}

func (p *PipeWrap) Exec(s string) *script.Pipe {
	if p == nil || p.Error() != nil {
		return p.Pipe
	}
	q := NewPipeWrap()
	//args := strings.Fields(s)
	cmd := exec.Command("bash", "-c", s)
	cmd.Stdin = p.Reader
	output, err := cmd.CombinedOutput()
	if err != nil {
		q.SetError(err)
	}
	return q.WithReader(bytes.NewReader(output))
}
func NewPipeWrap() *PipeWrap {
	PipeWrap := new(PipeWrap)
	PipeWrap.Pipe = script.NewPipe()
	return PipeWrap
}

func Exec(s string) *script.Pipe {
	return NewPipeWrap().Exec(s)
}
