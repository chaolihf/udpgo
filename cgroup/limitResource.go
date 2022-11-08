package cgroup

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/containerd/cgroups"
	"github.com/opencontainers/runtime-spec/specs-go"
)

const (
	MB = 1024 * 1024
)

func main() {
	shares := uint64(100)
	memoryLimit := int64(10000000)
	control, err := cgroups.New(cgroups.V1, cgroups.StaticPath("/test"), &specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Shares: &shares,
		},
		Memory: &specs.LinuxMemory{
			Limit: &memoryLimit,
		},
		Pids: &specs.LinuxPids{
			Limit: int64(os.Getpid()),
		},
	})
	pid := os.Getpid()
	control.Add(cgroups.Process{Pid: pid}, cgroups.Name("/test"))
	defer control.Delete()
	if err != nil {
		panic(err)
	}
	blocks := make([][MB]byte, 0)
	fmt.Println("Child pid is", os.Getpid())

	for i := 0; ; i++ {
		blocks = append(blocks, [MB]byte{})
		printMemUsage()
		time.Sleep(time.Second)
	}
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tSys = %v MiB \n", bToMb(m.Sys))
}

func bToMb(b uint64) uint64 {
	return b / MB
}
