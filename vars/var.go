package main

import "fmt"

const s string = "constant"
func main()  {
	
	var a string //declare a variable: string with a default value each type gives (eg - int: -1)
	var x,y string //declare multiple vars

	var c = "c" // declare and initialise string val
	var o,p int = 1,2
	var j,k = "hello", "world"
	l,m := 4,"test"
	d := "d" // shorter way to declare and init val : only works within a function

	fmt.Println(a,x,y,c,o,p,j,k,l,m,d)

	//constants
	const n = 50000
	// immutabe: n = 12
	var test_int = 1200
	// this works test_int = 22
	fmt.Println(s,n,test_int)

}