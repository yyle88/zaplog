package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetValuePointer(t *testing.T) {
	type exampleType struct {
		v string
	}
	a := exampleType{v: "a"}
	p := GetValuePointer(a)
	require.Equal(t, "a", p.v)
}

func TestGetPointerValue(t *testing.T) {
	type exampleType struct {
		v string
	}
	p := &exampleType{v: "a"}
	a := GetPointerValue(p)
	require.Equal(t, "a", a.v)
}

func TestPathUnescape(t *testing.T) {
	raw := "github.com/yyle88/zaplog/internal/examples/example1x/ZLG%e6%b5%8b%e9%9d%9eASCII%e8%b7%af%e5%be%84.TestZapLog"
	//goland:noinspection GoPrintFunctions
	t.Log(raw)
	res := PathUnescape(raw)
	t.Log(res)
}
