package parser

import (
	"bytes"
	"fmt"
	"sync/atomic"
)

var verbose = false

func SetVerbose(is bool) {
	verbose = is
}

var depth = atomic.Int32{}

// trace is a helper function for debugging.

const LOG_IDENT = "   "

func trace(name string) string {
	if !verbose {
		return name
	}
	depth.Add(1)

	printLog("BEGIN", " %s", name)

	return name
}

func getIdent() string {
	buf := bytes.Buffer{}
	for i := int32(0); i < depth.Load(); i++ {
		buf.WriteString(LOG_IDENT)
	}
	return buf.String()
}

func untrace(name string) {
	if !verbose {
		return
	}

	printLog("END", "   %s", name)

	depth.Add(-1)

	return
}

func printLog(sign, msg string, args ...interface{}) {
	fmt.Printf(getIdent()+signColor(sign)+msg+"\n", args...)
}

func signColor(s string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", s)
}
