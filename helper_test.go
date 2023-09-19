package jcli_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shipengqi/jcli"
)

func TestNormalizeCliName(t *testing.T) {
	name := jcli.NormalizeCliName("")
	assert.NotEmpty(t, name)

	name = jcli.NormalizeCliName("testname")
	assert.Contains(t, name, "testname")
}
