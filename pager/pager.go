package pager

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

// getUserPager retrieves the user's preferred pager from the PAGER environment variable or defaults to "less"
func getUserPager() []string {
	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "less -FRX"
	}
	return strings.Fields(pager)
}

type Pager struct {
	pageCmd *exec.Cmd
}

// NewPager creates a new Pager
func NewPager(stdout io.Writer, stderr io.Writer) *Pager {
	pager := getUserPager()
	cmd := exec.Command(pager[0], pager[1:]...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return &Pager{
		pageCmd: cmd,
	}
}

func (p *Pager) Page(content string) error {
	p.pageCmd.Stdin = strings.NewReader(content)
	return p.pageCmd.Run()
}
