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


### Why the declaration is different from C ?
Go believe their declaration style makes the complex declaration more readible.

### Reference
1. [A tour of Go](https://tour.golang.org/basics)
2. [Slice and Array](https://blog.golang.org/go-slices-usage-and-internals)
3. [Go中文筆記](https://openhome.cc/Gossip/Go/index.html)
