package main

import (
    "io"
    "strings"
    "fmt"
)

func rot13(b byte) byte {
    var a, z byte
    switch {
    case 'a' <= b && b <= 'z':
        a, z = 'a', 'z'
    case 'A' <= b && b <= 'Z':
        a, z = 'A', 'Z'
    default:
        return b
    }
    return (b-a+13)%(z-a+1) + a
}

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(p []byte) (int, error) {
    n, err := r.r.Read(p)
    for i := 0; i < n; i++ {
        p[i] = rot13(p[i])
    }
    return n, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    p := make([]byte, 20)
    r.Read(p)
    fmt.Print(p)
}
