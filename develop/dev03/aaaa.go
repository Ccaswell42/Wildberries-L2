package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sort", "stroki")
	val, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(val))
}
