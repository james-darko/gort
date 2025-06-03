package gort

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func Assert(b bool) {
	if Assertions && !b {
		log.Println("assertion failed")
		stack := debug.Stack()
		fmt.Println(string(stack))
		os.Exit(1)
	}
}

func Assertf(b bool, format string, args ...any) {
	if Assertions && !b {
		m := fmt.Sprintf(format, args...)
		log.Println(m)
		stack := debug.Stack()
		fmt.Println(string(stack))
		os.Exit(1)
	}
}

func AssertNoErr(err error) {
	if err != nil {
		log.Printf("got asserted error: %v\n", err)
		stack := debug.Stack()
		fmt.Println(string(stack))
		os.Exit(1)
	}
}

func AssertNoErrf(err error, format string, args ...any) {
	if err != nil {
		log.Printf(format+"\n", args...)
		stack := debug.Stack()
		fmt.Println(string(stack))
		os.Exit(1)
	}
}
