## Concurrency
### Why go routine is good?
1. [Go routine is not light weighr process](https://codeburst.io/why-goroutines-are-not-lightweight-threads-7c460c1f155f)
    Goroutine exitsts in only in the virtual space of go routine, not in the OS (different from thread). Thus, blocking is fine, because we can use channels in go which exactly only in virtual space, the OS doesn't block the thread. Such goroutine simply go in waiting state and other runnable goroutine (from M struct) is scheduled.
    Although threads are in the same virtual space, CPU still needs to save and flush many register, thus there is still overhead.
2. [Five things that make go fast](https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast) 
    1. Variable
    2. Inline
    3. Garbage collection
    4. Escape
    5. Goroutine

### Goroutine
A goroutine is a light weight thread managed by Go runtime.
```go
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}

```
If we want to use goroutine insiede a for loop, we should do like this
```go
for _, val := range values {
	go func(val interface{}) {
		fmt.Println(val)
	}(val)
}
```
If we don't pass the `val` to the inside function, the goroutine might not be executed until the for loop ends. Thus, we might see all goroutine print the last element. 
<br />
**Reference** <br />
1. [Closure and goroutine](https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables)

### Channels
```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
By deafult, send and receive block until the other side is ready. <br />

**Buffered Channel** <br />
```go
func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

**Range and Channel** <br />
Usually we don't need to close channel, only when receiver must be told that that there are no more values coming, such as closing a range.
```go
func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}
```
### Select
Select lets a goroutine wait on multiple communication operations.
```go
func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}
```

### Equivalent binary tree
```go
package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

func WalkTree(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        WalkTree(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        WalkTree(t.Right, ch)
    }
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    WalkTree(t, ch)
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1, ch2 := make(chan int), make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for i := range ch1 {
        if i != <-ch2 {
            return false
        }
    }
    return true
}

func main() {
    ch := make(chan int)
    go Walk(tree.New(1), ch)
    for i := range ch {
        fmt.Println(i)
    }
    fmt.Println("Should return true:", Same(tree.New(1), tree.New(1)))
    fmt.Println("Should return false:", Same(tree.New(1), tree.New(2)))
}
```

### sync.WaitGroup
Add would add one tasks to queue, done would remove one, and wait would force the main process to wait until all goroutines are done.
```go
var waitgroup sync.WaitGroup
 
func test(shownum int) {
	fmt.Println(shownum)
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}
 
func main() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go test(i)
	}
	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	fmt.Println("done!")
}
```
