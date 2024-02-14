// go run [filename]
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getStringInput (texto string) string {
	fmt.Print(texto)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.ReplaceAll(input, "'", "''")
	return input
}

func main(){
	var db []livro
	fmt.Println(db)
}
