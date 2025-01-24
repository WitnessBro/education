package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := getBuffers()
	defer out.Flush()

	var first, second int
	fmt.Fscan(in, &first)
	fmt.Fscan(in, &second)

	fmt.Fprintln(out, first+second)
}

func getBuffers() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	return in, out
}
