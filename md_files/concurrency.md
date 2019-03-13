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
