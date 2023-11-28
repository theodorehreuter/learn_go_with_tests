package main

import (
	"os"
	"time"

	"github.com/theodorehreuter/learn_go_with_tests/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
