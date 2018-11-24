// A stupid redis clone that doesn't do much

// Connect:
//    nc localhost 6379
//
// Commands:
//    EXIT        - end client session
//    SET foo bar - store value `bar` at key `foo`
//    GET foo     - get the value stored at `foo`
//
// That's all it does, told you it was stupid.

package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	var (
		err  error
		port = 6379
	)
	if len(os.Args) > 1 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Wanted port, got", os.Args[1])
		}
	}
	d := NewDatastore()
	h := NewHandler(d)
	s := NewServer(port, h)
	s.Serve()
}
