package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	name := os.Args[1]
	output := os.Args[2]

	src, err := ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	program := compile(src)

	fmt.Println(program)

	if err := WriteFile(output+".go", program); err != nil {
		fmt.Println("could not write file")
		return
	}

	_, err = exec.Command("go", "build", output+".go").Output()

	if err != nil {
		fmt.Println("could not compile")
		return
	}
}

func scan(str string, start int, char byte) int {
	count := 0
	for i := start; i < len(str); i++ {
		if str[i] != char {
			return count
		}
		count++
	}
	return count
}

func clean(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	return str
}

func compile(src string) string {
	src = clean(src)
	program := `package main
import "fmt"
func main() {
var mem [30000]byte
dp := 0
`
	for i := 0; i < len(src); i++ {
		switch src[i] {
		case '>':
			n := scan(src, i, '>')
			program += "dp += " + strconv.Itoa(n)
			i += n
		case '<':
			n := scan(src, i, '<')
			program += "dp -= " + strconv.Itoa(n)
			i += n
		case '+':
			n := scan(src, i, '+')
			program += "mem[dp] += " + strconv.Itoa(n)
			i += n
		case '-':
			n := scan(src, i, '-')
			program += "mem[dp] -= " + strconv.Itoa(n)
			i += n
		case '.':
			program += "fmt.Printf(\"%c\", mem[dp])"
		case ',':
			panic("input reading not implemented yet: line " + strconv.Itoa(i))
		case '[':
			program += "for mem[dp] != 0 {"
		case ']':
			program += "}"
		}
		program += "\n"
	}
	program += "}"
	return program
}

func ReadFile(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func WriteFile(path, text string) error {
	err := ioutil.WriteFile(path, []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
}
