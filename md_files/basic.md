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

When two or more consecutive named function parameters share the same type, we can omit all but last.
```go
func add(x, y int) int {
	return x + y
}
```

A function can return multiple result.
```go
func swap(x, y string) (string, string) {
	return y, x
}
```

Return values can be named like :
```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```
A return statement without arguments return the named return values. This is called naked return, it should only appear in short function.








