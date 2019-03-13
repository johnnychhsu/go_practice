package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    var s [][]uint8
    for i:=0; i<dy; i++ {
        tmp := make([]uint8, dx)
        s = append(s, tmp)
        for x:=0; x<dx; x++ {
            s[i][x] = uint8(uint8(x*i) / 2)
        }
    }
    return s
}

func main() {
    pic.Show(Pic)
}
