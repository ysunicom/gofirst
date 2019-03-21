package main

import "log"

type Connlimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *Connlimiter {
	return &Connlimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *Connlimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *Connlimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connection coming :%d", c)
}
