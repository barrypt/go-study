package main
 
import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"runtime"
)
 
// Example_workerPool demonstrates how to use a semaphore to limit the number of
// goroutines working on parallel tasks.
//
// This use of a semaphore mimics a typical “worker pool” pattern, but without
// the need to explicitly shut down idle workers when the work is done.
func main() {
	ctx := context.TODO()
 
	var (
		maxWorkers = runtime.GOMAXPROCS(0)
		sem        = semaphore.NewWeighted(int64(maxWorkers))
		out        = make([]int, 32)
	)
 
 
	// Compute the output using up to maxWorkers goroutines at a time.
	for i := range out {
		// When maxWorkers goroutines are in flight, Acquire blocks until one of the
		// workers finishes.
		if err := sem.Acquire(ctx, 1); err != nil { //获取锁、如果是Acquire会等待，而TryAcquire不会等待
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}
 
		go func(i int) {
			defer sem.Release(1)
			out[i] = collatzSteps(i + 1) //这个函数，只是一个举例子，不用太纠结
			//计算collatz这个函数到1的步骤，实际工作中是换成自己的任务
			// collatz函数             f(n)=  n/2 偶数
			//                               3*n+1 奇数
		}(i)
	}
 
	// Acquire all of the tokens to wait for any remaining workers to finish.
	//
	// If you are already waiting for the workers by some other means (such as an
	// errgroup.Group), you can omit this final Acquire call.
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}
 
	fmt.Println(out)
 
}
 
// collatzSteps computes the number of steps to reach 1 under the Collatz
// conjecture. (See https://en.wikipedia.org/wiki/Collatz_conjecture.)
func collatzSteps(n int) (steps int) {
	if n <= 0 {
		panic("nonpositive input")
	}
 
	for ; n > 1; steps++ {
		if steps < 0 {
			panic("too many steps")
		}
 
		if n%2 == 0 {
			n /= 2
			continue
		}
 
		const maxInt = int(^uint(0) >> 1)
		if n > (maxInt-1)/3 {
			panic("overflow")
		}
		n = 3*n + 1
	}
 
	return steps
}