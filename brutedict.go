package brutedict

import (
	"bytes"
)

type BruteDict struct {
	isnum   bool
	islow   bool
	iscap   bool
	start   int
	end     int
	queue   chan string
	quit    chan bool
	running bool
}

func New(isnum, islow, iscap bool, start, end int) (bd *BruteDict) {
	bd = &BruteDict{
		isnum:   isnum,
		islow:   islow,
		iscap:   iscap,
		start:   start,
		end:     end,
		running: true,
		queue:   make(chan string),
		quit:    make(chan bool),
	}

	strnum := []byte("0123456789")
	strlow := []byte("abcdefghijklmnopqrstuvwxyz")
	strcap := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var str = make([]byte, 0)
	if isnum {
		str = append(str, strnum...)
	}
	if islow {
		str = append(str, strlow...)
	}
	if iscap {
		str = append(str, strcap...)
	}

	var b = make([]byte, end)
	go bd.process(str, b, start, end)
	return
}

func (bd *BruteDict) process(str []byte, b []byte, start int, end int) {
	defer func() { recover() }()

	for i := start; i <= end; i++ {
		bd.list(str, b, i, 0)
	}
	bd.quit <- true
}

func (bd *BruteDict) Id() (str string) {
	select {
	case str = <-bd.queue:
	case <-bd.quit:
	}
	return
}

func (bd *BruteDict) Close() {
	bd.running = false
	close(bd.queue)
}

func (bd *BruteDict) list(str []byte, b []byte, l int, j int) {
	strl := len(str)

	for i := 0; i < strl; i++ {
		b[j] = str[i]
		if j+1 < l {
			bd.list(str, b, l, j+1)
		} else {
			n := bytes.IndexByte(b, 0)
			bd.queue <- string(b[:n])
		}
	}
}
