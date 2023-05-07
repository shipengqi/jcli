package jcli

import "github.com/fatih/color"

func Red(msg string) string {
	return color.RedString(msg)
}

func Yellow(msg string) string {
	return color.YellowString(msg)
}

func Green(msg string) string {
	return color.GreenString(msg)
}

func Blue(msg string) string {
	return color.BlueString(msg)
}

func Colorize(msg string, attrs ...color.Attribute) string {
	co := color.New(attrs...)
	return co.Sprint(msg)
}
