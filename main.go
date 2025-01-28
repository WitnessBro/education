package main

import (
	_ "github.com/lib/pq"

	"github.com/WitnessBro/education/cmd"
)

func main() {
	cmd.Execute()
}
