package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Buscar: ")
	query, _ := reader.ReadString('\n')
	query = strings.Trim(query, "\n")

	inSpain, err := Search(query)
	if err != nil {
		fmt.Println(err)
		return //print error...
	}

	fmt.Println("En españa es: " + inSpain)
}
