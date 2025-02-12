package logger

import (
	"strconv"
	"strings"
	"testing"
)

func Test_rpcid(t *testing.T) {

	extra := "1"
	last := strings.LastIndex(extra, ".")
	i, _ := strconv.Atoi(extra[last+1:])
	extra = extra[:last+1] + strconv.Itoa(i+1)
	t.Log(extra)
}
