package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// register statement on defer
	i := 0
	defer fmt.Println("defer function 1, value of i is", i)
	defer fmt.Println("defer function 2")
	defer fmt.Println("defer function 3")
	i++
	fmt.Println("The current value of i is", i)

	// example: file object close
	// get current directory path
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	f, _ := ReadFile(path +"/input.txt")
	fmt.Println(string(f))

	// function defer
	defer who()
	hello()
}


func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func hello() {
	fmt.Println("Hello")
}

func who() {
	fmt.Println("Go")
}