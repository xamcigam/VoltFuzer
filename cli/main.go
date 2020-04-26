package main

import (
	"fmt"
	"github.com/xamcigam/VoltFuzer"
)

func main() {
	var master *Series
	var para *Parallel
	var w1, w2, w3

	master.childA = w1
	w1.parent = master
	master.childB = para
	para.parent = master
	para.childA = w2
	w2.parent = para
	para.childB = w3
	w3.parent = para

	fmt.Println(master.IsValid())
}
