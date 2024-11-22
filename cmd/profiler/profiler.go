package profiler

import (
	"os"
	"runtime/pprof"

	"gobox/internal/executor"
)

func main() {
	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	executor.ExecuteSolution(executor.SolutionCase)
}
