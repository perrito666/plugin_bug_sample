package main

import "plugin"

func main() {
	p, err := plugin.Open("./v01/v0.1.so")
	//p, err := plugin.Open("./v01/v01.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"
}
