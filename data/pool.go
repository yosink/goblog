package data

import (
	"github.com/pkg/errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size < 0 {
		return nil, errors.New("Size value was too small")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}
func (pool *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-pool.resources:
		log.Println("Acquire:", "shared resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "new resource")
		return pool.factory()

	}

}
func (pool *Pool) Release(r io.Closer) {
	pool.m.Lock()
	defer pool.m.Unlock()
	if pool.closed {
		r.Close()
		return
	}
	select {
	case pool.resources <- r:
		log.Println("Release:", "in queue")
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
