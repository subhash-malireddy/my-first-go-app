package helper

import "github.com/fatih/color"

func Greet() *color.Color {
	return color.New(color.FgBlue, color.Bold)
}

func Info() *color.Color {
	return color.New(color.FgYellow)
}

func Error() *color.Color {
	return color.New(color.FgRed, color.Bold)
}

func Success() *color.Color {
	return color.New(color.FgGreen, color.Bold)
}

func Question() *color.Color {
	return color.New(color.FgCyan, color.Bold)
}
func Input() *color.Color {
	return color.New(color.FgHiMagenta)
}
