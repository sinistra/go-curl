package main

import (
	"time"

	"github.com/rs/zerolog"

	"go-curl/cmd"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	cmd.Execute()
}
