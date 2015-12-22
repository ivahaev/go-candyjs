package main

import (
	"os"

	"github.com/ivahaev/go-candyjs"
)

//go:generate candyjs import net
//go:generate candyjs import net/http
//go:generate candyjs import io/ioutil
func main() {
	script := os.Args[1]
	ctx := candyjs.NewContext()
	ctx.PevalFile(script)
}
