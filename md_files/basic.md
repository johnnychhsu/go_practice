## Basic
### Functions
Notice the type comes after the variable names
```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```
<br />

When two or more consecutive named function parameters share the same type, we can omit all but last.
```go
func add(x, y int) int {
	return x + y
}
```
<br />

A function can return multiple result.
```go
func swap(x, y string) (string, string) {
	return y, x
}
```
<br />

Return values can be named like :
```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```
A return statement without arguments return the named return values. This is called naked return, it should only appear in short function.
<br />

### Variables
The var statement declare a list of variables
```go
var c, python, java bool
```
<br />

A var can include initializer, one per variable. If an initializer is present, the type can be omitted.
```go
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```
<br />

Inside a function, the `:=` short assignment statement can be used in place of a var with implicit type.
```go
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```
<br />

Variables can be factor into blocks like import 
```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```
<br />

Variables declared without an explicit initial value are given their zero value.
```go
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
// 0 0 false ""
```
<br />

Type conversion
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// can also be
i := 42
f := float64(i)
u := uint(f)
```
<br />
Type inference, if the type is not explicit, the type is inferenced from the value assigned to it.
```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```
<br />

Constant, can't not be declared with `:=`.
```go
const World = "世界"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
// Big will overflow
```

### Basic Type
Go's basic type are
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
The `int`, `uint` and `uintptr` types usually use 32 bits on a 32 bits system, 64 bits on a 64 bits system. We should use `int` for int type unless we have specific reason.
<br />

### For loop
Go only have for loop. Don't need `()` but `{}` is required.
```go
for i := 0; i < 10; i++ {
		sum += i
	}
// the below is while concept
for sum < 1000 {
		sum += sum
	}
```

### If
`if` statement can start with a short statement to execute before the condition. Variable declared here can only be used till the end of `if`.
```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return 0
}
```

### Switch
Go's switch only runs the selected case, the rest will not be executed. The switch case need not be constant.
```go
switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
```
<br />

Switch without consition is `switch true`.
```go
switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
```

### Defer
A defer function defers the execution of a function until the surrouding function returns.
```go
defer fmt.Println("world")
```
<br />

Defer stack, defered function calls are pushed onto a stack.

<<<<<<< HEAD
### Struct
```go
type Vertex struct {
	X int
	Y int
}

v := Vertex{1, 2}
v.X := 7
p := &v
p.X = 1e9

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
	
)
```

### Struct 
```go
type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}  // has type Vertex
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X:0 and Y:0
    p  = &Vertex{1, 2} // has type *Vertex
)
```

### Array
An array's length is part of its type, so arrays can't be resized.
```go
var a [2]string
a[0] = "Hello"
a[1] = "World"

primes := [6]int{2, 3, 5, 7, 11, 13}
```

### Slices
It is dynamically sized. It is formed by soecifying two indices, lower and upper bound. This will contain the first element, but excludes the last element.
A slice is a struct contain : 1
1. A pointer to the head of under-lying array.
2. length of the slice.
3. capacity of the under-lying array.
<br />
Slices are like reference of array. Change the content of slice will change the array accordingly. <br />
A slice literal is like array literal without the length. <br />
We can omit the two bounds, the default is `[0: length_of_slice]`.
```go
// This is array
primes := [6]int{2, 3, 5, 7, 11, 13}

// This is slice
var s []int = primes[1:4]

// slice of type struct
s := []struct {
        i int
        b bool
}{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
}
```
<br />

**Slice length and capacity** <br />
The length of the slice is the number of elements it contains. <br />
The capacity of a slice is the number of elements in the underlying array.
```go
s := []int{2, 3, 5, 7, 11, 13}
printSlice(s)

// Slice the slice to give it zero length.
s = s[:0]
printSlice(s)

// Extend its length.
s = s[:4]
printSlice(s)

// Drop its first two values, can't bring back those dropped
s = s[2:]
printSlice(s)

// print
// len=6 cap=6 [2 3 5 7 11 13]
// len=0 cap=6 []
// len=4 cap=6 [2 3 5 7]
// len=2 cap=4 [5 7]

```
The zero value of a slice is `nil`.
```go
var s []int
s == nil
```

**Create slice with make and copy** <br />
```go
a := make([]int, 5)
printSlice("a", a)

b := make([]int, 0, 5)
printSlice("b", b)

c := b[:2]
printSlice("c", c)

d := c[2:5]
printSlice("d", d)

// a len=5 cap=5 [0 0 0 0 0]
// b len=0 cap=5 []
// c len=2 cap=5 [0 0]
// d len=3 cap=3 [0 0 0]
```

**Slice of slices** <br />
```go
// Create a tic-tac-toe board.
board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
}
```

**Append** <br />
```go
var a []int
a = append(a, 1, 2, 3)
```

**A possible usage** <br />
We only need the specific value in the entire file, but because we have a slice reference to it, the entire array can't not br free. Thus we can copy those we are interested in, then return the new slice.
```go
func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}
```

### Range
It is like `enumerate` in python, `range` return index and element in the slice.
```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}
```

### Map
Like dictionary in python.
```go
m := make(map[string]Vertex)
m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
}

var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}

// We can omit the type
var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}
```

**Mutating Map** <br />
We can use `ele, ok` to get element. If the key is not there, `ok` will be `False`.
```go
func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])

    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])

    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
```
### Function closure
We can't define a function inside a function in Go. But we can pass a function as parameter into a function. If a function use a variable outside its scope, then we say there is a closure. The inner function will keep the variable itself, not only the value (like a pointer),
```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}
```
