package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "I am a good man, he is a good man, too."
    words := strings.Fields(s)
    ans := make(map[string]int)
    for _, word := range words {
        if v, ok := ans[word]; ok {
            ans[word] = v + 1
        } else {
            ans[word] = 1
        }
    }
    fmt.Print(ans)
}

