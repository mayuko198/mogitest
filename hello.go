package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//file open
	f, err := os.Open("Hello.txt")

	if err != nil {
		fmt.Println("error")
	}

	defer f.Close()

	//read
	b, err := ioutil.ReadAll(f)

	//output
	fmt.Println(string(b))
}
