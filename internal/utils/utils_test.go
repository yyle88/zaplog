package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPtrX(t *testing.T) {
	type aType struct {
		v string
	}
	a := aType{v: "a"}
	p := PtrX(a)
	require.Equal(t, "a", p.v)
}

func TestSoftPathUnescape(t *testing.T) {
	raw := "github.com/yyle88/zaplog/internal/examples/example1x/ZLG%e6%b5%8b%e9%9d%9eASCII%e8%b7%af%e5%be%84.TestZapLog"
	t.Log(raw)
	res := SoftPathUnescape(raw)
	t.Log(res)
}
