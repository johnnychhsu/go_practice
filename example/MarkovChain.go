package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Prefix []string

func (p Prefix) String() string {
	return strings.Join(p, "")
}

func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain struct {
	chain     map[string][]string
	prefixLen int
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != err {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var word []string
	for i := 0; i < n; i++ {
		choice := c.chain[p.String()]
		if len(choice) == 0 {
			break
		}
		next := choice[rand.Intn(len(choice))]
		word = append(word, next)
		p.Shift(next)
	}
	return strings.Join(word, " ")
}

func main() {
	numWords := flag.Int("words", 100, "Maximum number if words to print")
	prefixLen := flag.Int("prefix", 2, "Prefix length in words")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	c := NewChain(*prefixLen)
	c.Build(os.Stdin)
	text := c.Generate(*numWords)
	fmt.Println(text)
}
