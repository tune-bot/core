package core

import (
	"fmt"
	"os"
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

var debugEnv, isSet = os.LookupEnv("DEBUG")

var isDebugOverride = false // set this to true to ignore environment variable
var isDebug = (isSet && debugEnv != "false" && debugEnv != "0" && debugEnv != "") || isDebugOverride

var debugColourInRotation = yellow

func rotageDebugColor() color {
	switch debugColourInRotation {
	case yellow:
		debugColourInRotation = blue
	case blue:
		debugColourInRotation = yellow
	}
	return debugColourInRotation
}

var successColourInRotation color = cyan

func rotateSuccessColor() color {
	switch successColourInRotation {
	case green:
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
		errorColourInRotation = magenta
	case magenta:
		errorColourInRotation = red
	}
	return errorColourInRotation
}

func PrintLnColor(msg string, colour color) {
	fmt.Printf("\033[%dm%s\033[0m\n", colour, msg)
}

func PrintDebug(msg string) {
	var a, b = os.LookupEnv("DEBUG")
	PrintSuccess(a)
	PrintSuccess(fmt.Sprintf("%t", b))

	if isDebug {
		PrintLnColor(msg, rotageDebugColor())
	}
}

func PrintError(msg string) {
	PrintLnColor(msg, rotateErrorColor())
}

func PrintSuccess(msg string) {
	PrintLnColor(msg, rotateSuccessColor())
}
