package pager

import (
	"os"
	"testing"
)

func TestPager_Page(t *testing.T) {
	pager := NewPager(os.Stdout, os.Stderr)
	_ = pager.Page("salkdjf\nasdf\n\n\n\n\n\n\n\n\nlaksjdflaksjdflk\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nasjdf")
}
