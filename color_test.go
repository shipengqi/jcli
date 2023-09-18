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
	str := jcli.Yellow("yellow string")
	t.Log(str)
	assert.Equal(t, str, "yellow string")
}

func TestGreen(t *testing.T) {
	str := jcli.Green("green string")
	assert.Equal(t, str, "green string")
}

func TestBlue(t *testing.T) {
	str := jcli.Blue("blue string")
	assert.Equal(t, str, "blue string")
}
