package progressBar

import (
	"testing"
	"time"
	"fmt"
	"strconv"
	"sync"
)

func TestProgress_Bar(t *testing.T) {
	var nb = NewBar(100, MODEL_NUMBER, "progress: ", "!!!!", false)
	nb.Start()

	//var c = make(chan string)
	//go func() {
	//	c <- "sh"
	//}()
	//
	//var a = <-c
	//fmt.Println(a)
	//go func() {
	//	time.Sleep(time.Second)
	//	nb.c <- "1"
	//}()
	for i := 0; i < 100; i++ {
		nb.Count(1)
		time.Sleep(10 * time.Millisecond)
	}
	nb.Wait()
	return

	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go func(syncWg *sync.WaitGroup) {
		for {
			nb.shower.Show(nb.Current, nb.Total, nb.Prefix, nb.Suffix, nb.isPercentage)
			if nb.Current >= nb.Total {
				break
			}

			var current = <-nb.c
			nb.Current, _ = strconv.Atoi(current)
			fmt.Println("-----------", nb.Current, nb.Total, nb.Prefix)
		}
		syncWg.Done()
	}(wg)
	wg.Wait()

	return

	//for i := 0; i < 100; i++ {
	//	nb.Count(1)
	//	time.Sleep(500 * time.Millisecond)
	//}

	time.Sleep(1)
}