package main

import 	(
	witlog "go.wit.com/log"
)

// various debugging flags
var logNow bool = true	// useful for active development
var logError bool = true
var logWarn bool = false
var logInfo bool = false
var logVerbose bool = false

var outputS []string

func log(b bool, a ...any) {
	witlog.Log(b, a...)
}

func sleep(a ...any) {
	witlog.Sleep(a...)
}

func exit(a ...any) {
	witlog.Exit(a...)
}
