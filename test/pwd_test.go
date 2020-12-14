package test

import (
	u "go-admin/pkg/util"
	"testing"
)

func TestMD5(t *testing.T) {
	t.Log(u.EncodeMD5("test"))
}
