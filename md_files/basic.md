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

