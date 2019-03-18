## My Go practice code
### Go Lang notes
1. [Basic syntax](md_files/basic.md)
2. [Concurrency](md_files/concurrency.md)

### Some special feature in Go
**Defer** <br />
A defer statement will hold until surrounding function finish executing.

### Function closure
Fibonacci generator using function closure.
Reference : https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/

### Function Receiver
No matter the receiver receive value or pointer, both value and pointer can use it.
We prefer receiver type is pointer due to some reasons:
1. The method can modify the value that its receiver points to.
2. Avoid copying the value on each method call. This can be more efficient if the receiver is a large struct.

### Interface
It is like a top level representation, different type implement it, as the interface's initialization.

### Defer
Three rules : 
1. A deferred function's arguments are evaluated when the defer statement is evaluated.
The deferred call will print 0.
```go
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```
2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
It will print 3210.
```go
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
```
3. Deferred functions may read and assign to the returning function's named return values.

### Why the declaration is different from C ?
Go believe their declaration style makes the complex declaration more readible.

### Projects
1. [A simple server example](https://github.com/johnnychhsu/wikigo)

### Reference
1. [A tour of Go](https://tour.golang.org/basics)
2. [Slice and Array](https://blog.golang.org/go-slices-usage-and-internals)
3. [Go中文筆記](https://openhome.cc/Gossip/Go/index.html)
4. [Defer](https://blog.golang.org/defer-panic-and-recover)
