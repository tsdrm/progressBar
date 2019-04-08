# progress
A simple program implement the multi progress bar on shell.

### Progress Bar
progressBar is to support a progress bar displayed on the shell, support a single progress bar, also supports multiple progress bars, running on windows will be abnormal, because the AnsiEscapeCodes is not supported in the windows part, such as \033 [NA, etc

#### Basic Usage

##### Single Bar

```go
package main

import (
	pb "github.com/tsdrm/progressBar"
	"time"
)

func main() {
    var total = 123
	var b = pb.NewBar(total, pb.MODEL_NUMBER, "progress: [", "] !!!", false)
	b.Start()

	go func() {
		for i := 0; i < total; i ++ {
			b.Count(1)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	time.Sleep(time.Millisecond * 5)
	b.Wait()
}
```

##### Multi Bar
```go
package main

import (
    pb "github.com/tsdrm/progressBar"
    "time"
)
var processRun = func(p *pb.Progress, count int, n int, interval time.Duration) {
	for i := 0; i < count; i++ {
		p.Count(n)
		time.Sleep(interval)
	}
}

func main()  {
    var count = 145
    var pg = pb.NewProcessGroup()
    var p1 = pb.NewBar(count, pb.MODEL_NUMBER, "progressA: [", "]!!!", true)
    var p2 = pb.NewBar(count, pb.MODEL_NUMBER, "progressB: [", "]...", false)
    pg.Add(p1)
    pg.Add(p2)
    pg.Start()
    go processRun(p1, count, 1, time.Millisecond*10)
    go processRun(p2, count, 1, time.Millisecond*10)
    time.Sleep(10 * time.Millisecond)
    pg.Wait()
}

```