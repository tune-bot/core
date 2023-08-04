package core

import (
	"fmt"
)

type color int

var black color = 30
var red color = 31
var green color = 32
var yellow color = 33
var blue color = 34
var magenta color = 35
var cyan color = 36
var white color = 37
var reset color = 0

var successColourInRotation color = cyan

func rotateSuccessColor() color {
	switch successColourInRotation {
	case green:
		successColourInRotation = blue
	case blue:
		successColourInRotation = cyan
	case cyan:
		successColourInRotation = green
	}
	return successColourInRotation
}

var errorColourInRotation color = magenta

func rotateErrorColor() color {
	switch errorColourInRotation {
	case red:
		errorColourInRotation = yellow
	case yellow:
		errorColourInRotation = magenta
	case magenta:
		errorColourInRotation = red
	}
	return errorColourInRotation
}

func PrintLnColor(msg string, colour color) {
	fmt.Printf("\033[%dm%s\033[0m\n", colour, msg)
}

func PrintError(msg string) {
	PrintLnColor(msg, rotateErrorColor())
}

func PrintSuccess(msg string) {
	PrintLnColor(msg, rotateSuccessColor())
}
