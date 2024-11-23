package utils

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func Profile() func() {
	fmt.Println("Profiling...")

	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		panic(err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		f.Close()
		panic(err)
	}

	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}
}
