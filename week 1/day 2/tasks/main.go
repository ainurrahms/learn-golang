package main

import "fmt"

func calcLuasTabung() {
	T := 20.0
	r := 4.0
	phi := 3.14
	Lp := 2*(phi*(r*r)) + 2*(phi*(r*T))
	fmt.Println(Lp)
}

func main() {
	calcLuasTabung()
}
