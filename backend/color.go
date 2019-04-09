package main

// ANSI Escape colors
// https://stackoverflow.com/questions/6555995/ansi-escape-sequences-as-bytes
var (
	COLOR_GREEN = string([]byte{27, 91, 57, 55, 59, 51, 50, 59, 49, 109})
	COLOR_CYAN  = string([]byte{27, 91, 57, 55, 59, 51, 54, 59, 49, 109})
	COLOR_RESET = string([]byte{27, 91, 48, 109})
)
