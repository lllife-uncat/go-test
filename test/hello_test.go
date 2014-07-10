package test

import "testing"
import (
	"fmt"
)

type Cat struct {
	x string
	y string
}

func (c *Cat) eat() string {
	return "Jub jub"
}

func (c *Cat) walk() string {
	return "Jok jok"
}

func (c *Cat) run() string {
	return "Fas fas"
}

func TestCat(_ *testing.T) {
	var cat = new(Cat)
	var walk = cat.walk()
	var eat = cat.eat()
	var run = cat.run()
	
	fmt.Println(walk)
	fmt.Println(eat)
	fmt.Println(run)
}

func TestDoc(_ *testing.T) {
	fmt.Println("test doc...")
	fmt.Println("test doc...")
}
