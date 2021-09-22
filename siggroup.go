package siggroup

import (
	"io/ioutil"
	"os"
	"os/signal"
	"siggroup/errcause"
	"syscall"
)

// Quit Send process exit signal
func Quit() {
	sig <- syscall.SIGQUIT
}

// Add new Goroutine
func Add(routine func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				errcause.Keep(err)
			}
			Quit()
		}()
		routine()
	}()
}

// Wait Listen process exit signal
func Wait(cancel func()) {
	signal.Notify(sig, _signal...)
	for message := range sig {
		for i := range _signal {
			if message == _signal[i] {
				_ = ioutil.WriteFile("signal.txt", []byte(message.String()), 0666)
				if cancel != nil {
					cancel()
				}
				return
			}
		}
	}
}

var _signal = []os.Signal{
	syscall.SIGHUP,  //  1:  hangup
	syscall.SIGINT,  //  2:  interrupt
	syscall.SIGQUIT, //  3:  quit
	syscall.SIGILL,  //  4:  illegal instruction
	syscall.SIGTRAP, //  5:  trace/breakpoint trap
	syscall.SIGABRT, //  6:  aborted
	syscall.SIGBUS,  //  7:  bus error
	syscall.SIGFPE,  //  8:  floating point exception
	syscall.SIGSEGV, // 11:  segmentation fault
	syscall.SIGALRM, // 14:  alarm clock
	syscall.SIGTERM, // 15:  terminated
}

var sig = make(chan os.Signal)