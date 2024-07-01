There is a function that searches for something in a file. But it makes it not very quickly. We need to optimize it.

Assignment to work with the pprof profiler.

The goal of the task is to learn how to work with pprof, find hot spots in the code, be able to build a CPU and memory consumption profile, and optimize the code taking this information into account. Writing the fastest solution is not the goal of the assignment.

To generate the graph you will need graphviz. For windows users, be sure to add it to PATH so that the dot command is available.

I recommend that you read the addendum carefully. There are many more examples of optimization and explanations of how to work with the profiler. In fact, there is all the information to complete this task.

There are a dozen places where you can optimize.
You need to write a report about where you optimized and what. With screenshots and explanation of what they did. To learn exactly how to find problems in pprof, and not use your brains to figure out what’s slow here.

To complete the task, one of the parameters ( ns/op, B/op, allocs/op ) must be faster than in *BenchmarkSolution* ( fast < solution ) and another one is better than *BenchmarkSolution* + 20% ( fast < solution * 1.2) , for example ( fast allocs/op < 10422*1.2=12506 ).

In terms of memory ( B/op ) and the number of allocations ( allocs/op ), you can focus exactly on the results of *BenchmarkSolution* below, in terms of time ( ns/op ) - no, it depends on the system.

There is no need to parallelize (use goroutines) or sync.Pool for this task.

The result in fast.go into the FastSearch function (initially the same as in SlowSearch).

An example of the results that will be compared with:
```
$ go test -bench . -benchmem

goos: windows

goarch: amd64

BenchmarkSlow-8 10 142703250 ns/op 336887900 B/op 284175 allocs/op

BenchmarkSolution-8 500 2782432 ns/op 559910 B/op 10422 allocs/op

PASS

ok coursera/hw3 3.897s
```

Launch:
* `go test -v` - to check that nothing is broken
* `go test -bench . -benchmem` - to view performance
* `go tool pprof -http=:8083 /path/ho/bin /path/to/out` - web interface for pprof, use it to search for hot spots. Don't forget that you have 2 modes - cpu and mem, there are different out files.

Adviсe:
* See where we allocate memory
* See where we accumulate the entire result, although we don’t need all the values ​​at the same time
* See where type conversions occur that can be avoided
* Look not only at the graph, but also in pprof in text form (list FastSearch) - there you can see where everything is directly from the source code
* The task assumes the use of easyjson. This library is on the server, you can connect it. But you need to place the code generated via easyjson in a file with your function
* Can be done without easyjson

Note:
* easyjson is based on reflection and cannot work with the main package. To generate code, you need to put your structure into a separate package, generate the code there, then pick it up in main
