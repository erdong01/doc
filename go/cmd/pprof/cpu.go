package pprof

import (
	"log"
	"sync"
	"time"
)

type Cpu struct {
}

func (c *Cpu) Name() string {
	return "cpu"
}

func (c *Cpu) Run() {
	log.Println(c.Name(), "Run")
	for i := 0; i < 1000000000; i++ {
	}
}

type Block struct {
}

func (c *Block) Name() string {
	return "block"
}

func (c *Block) Run() {
	log.Println(c.Name(), "Run")
	<-time.After(time.Second)
}

type Goroutine struct {
}

func (g *Goroutine) Name() string {
	return "goroutine"
}

func (g *Goroutine) Run() {
	log.Println(g.Name(), "Run")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}

const (
	I = 1 << (iota * 10)
	Ki
	Mi
	Gi
	Ti
	Pi
)

type Mem struct {
	buffer [][Mi]byte
}

func (m *Mem) Name() string {
	return "mem"
}

func (m *Mem) Run() {
	log.Println(m.Name(), "Run")
	for len(m.buffer)*Mi < Gi {
		m.buffer = append(m.buffer, [Mi]byte{})
	}
}

type Mutex struct {
}

func (m *Mutex) Name() string {
	return "mutex"
}

func (m *Mutex) Run() {
	log.Println(m.Name(), "Run")
	mutex := &sync.Mutex{}
	mutex.Lock()
	go func() {
		time.Sleep(time.Second)
		mutex.Unlock()
	}()
	mutex.Lock()
}

type Cmd interface {
	Name() string
	Run()
}
