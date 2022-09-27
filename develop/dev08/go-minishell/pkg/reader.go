package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadlLine() {
	fmt.Println("Minishell starts")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", "MINISHELL$ ")
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			break
		}
		arr := ParseCmd(scanner.Text())
		err := ExecuteCmd(arr, os.Stdin, os.Stdout)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s", "MINISHELL$ ")
	}
}
