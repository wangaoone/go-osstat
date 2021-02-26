// +build linux darwin,cgo

package main

import (
	"github.com/wangaoone/go-osstat/cpu"
)

type cpuGenerator struct {
	cpu *cpu.Stats
	err error
}

func (gen *cpuGenerator) Get() {
	//gen.cpu, gen.err = cpu.Get(1)
	panic("implement me")
}

func (gen *cpuGenerator) Error() error {
	return gen.err
}

func (gen *cpuGenerator) Print(out chan<- value) {
	cpu := gen.cpu
	out <- value{"cpu.user", cpu.User, "-"}
	out <- value{"cpu.system", cpu.System, "-"}
	out <- value{"cpu.idle", cpu.Idle, "-"}
	out <- value{"cpu.total", cpu.Total, "-"}
}
