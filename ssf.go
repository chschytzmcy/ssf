package ssf

import (
	"os"
	"sync"

	"github.com/chschytzmcy/ssf/base/object"
)

type Cssf struct {
	object.Cobject
	wq      sync.WaitGroup
	sig     chan os.Singnal
	resetch chan int

	health int
}

func (s *Cssf) Init(name string) error {
	println("ssf init")
}

func (s *Cssf) Run() {
	println("ssf run")
}
