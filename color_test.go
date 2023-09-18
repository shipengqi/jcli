package jcli_test

import (
	"github.com/fatih/color"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shipengqi/jcli"
)

func TestRed(t *testing.T) {
	str := jcli.Red("red string")
	assert.Equal(t, str, "red string")
}

func TestYellow(t *testing.T) {
	str := jcli.Yellow("yellow string")
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

func TestIconBlue(t *testing.T) {
	str := jcli.IconBlue("icon blue string")
	assert.Equal(t, str, "icon blue string")
}

func TestColorize(t *testing.T) {
	str := jcli.Colorize("bg blue string", color.BgBlue)
	assert.Equal(t, str, "bg blue string")
}
