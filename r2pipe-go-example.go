package main

import (
	"fmt"

	"github.com/radare/r2pipe.go"
)

func main() {
	r2p, err := r2pipe.NewPipe("/bin/ls")
	if err != nil {
		panic(err)
	}
	defer r2p.Close()

	_, err = r2p.Cmd("aaaa")
	if err != nil {
		panic(err)
	}
	buf0, err := r2p.Cmd("S*")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf0)

	buf1, err := r2p.Cmd("pi 10")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf1)
	
	buf2, err := r2p.Cmd("px 64")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf2)
}
