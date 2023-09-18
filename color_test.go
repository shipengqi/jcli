package jcli_test

import (
	"github.com/shipengqi/jcli"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRed(t *testing.T) {
	str := jcli.Red("red string")
	assert.Equal(t, str, "red string")
}

func TestYellow(t *testing.T) {
	str := jcli.Yellow("red string")
	assert.Equal(t, str, "red string")
}

func TestGreen(t *testing.T) {
	str := jcli.Green("red string")
	assert.Equal(t, str, "red string")
}

func TestBlue(t *testing.T) {
	str := jcli.Blue("red string")
	assert.Equal(t, str, "red string")
}
