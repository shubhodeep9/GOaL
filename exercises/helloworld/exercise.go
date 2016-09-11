package helloworld

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Output() {
	gopath := os.Getenv("GOPATH")
	problem, err := ioutil.ReadFile(gopath + "/src/github.com/shubhodeep9/goal/exercises/helloworld/problem.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(problem))
}
