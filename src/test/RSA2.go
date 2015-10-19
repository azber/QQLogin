package main


import (
	"github.com/robertkrimen/otto"
	"fmt"
)

func main() {

	vm := otto.New()
	_,err := vm.Run(`
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