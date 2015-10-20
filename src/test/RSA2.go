package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

func main() {

	vm := otto.New()
	_, err := vm.Run(`
		function haha(str){
			console.log(str)
		}
	`)
	if err != nil {
		panic(err)
	}
	vm.Run(`
		haha("1234");
	`)
}
