package main

import (
	"flag"
	"fmt"
	"os"
	"quote-me/config"
	"quote-me/server"
)

func main() {
	env := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: quote-me -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*env)
	server.Init()
}