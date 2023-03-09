package pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	ff, _ := os.Create("mem.pprof")
	i := 0
	for {
		ICaller()
		if i > 200000 {
			break
		}
		i++
	}
	_ = pprof.WriteHeapProfile(ff)
}

var cacheBuff []string

func ICaller() {
	cacheBuff = append(cacheBuff, strings.Repeat("caller", 100))
	ISubCaller()
}

func ISubCaller() {

}

var cmds = []Cmd{
	&Cpu{},
	&Mem{},
	&Block{},
	&Goroutine{},
	&Mutex{},
}

func TestMain(t *testing.T) {
	log.SetFlags(log.Llongfile)
	log.SetOutput(os.Stdout)
	//开启对锁跟踪
	runtime.SetMutexProfileFraction(1)
	//开启对锁跟踪
	runtime.SetBlockProfileRate(1)
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	for {
		for _, v := range cmds {
			v.Run()
		}
		time.Sleep(time.Second)
	}
}
