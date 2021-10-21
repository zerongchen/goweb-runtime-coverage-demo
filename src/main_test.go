package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

var systemTest *bool
var exit *bool

func init() {
	systemTest = flag.Bool("systemTest", false, "Set to true when running system tests")
	exit = new(bool)

}

// Test started when the test binary is started. Only calls main.
func TestSystem(t *testing.T) {
	if !*systemTest {
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("ready to exit:", s)
				*exit = true
			default:
				fmt.Println("其他信号:", s)
			}
		}
	}()
	go func() {
		main()
	}()
	for !*exit {
		time.Sleep(5 * time.Second)
	}
}
