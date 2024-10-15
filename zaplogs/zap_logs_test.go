package zaplogs

import "testing"

func TestSoftUnescape(t *testing.T) {
	raw := "github.com/yyle88/zaplog/internal/examples/example1x/ZLG%e6%b5%8b%e9%9d%9eASCII%e8%b7%af%e5%be%84.TestZapLog"
	t.Log(raw)
	res := softUnescape(raw)
	t.Log(res)
}
