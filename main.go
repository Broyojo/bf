package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fileName := "mandelbrot.b"
	outputName := "hello"

	src, err := ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	program := transpile(src)

	if err := WriteFile(outputName+".go", program); err != nil {
		fmt.Println("could not write file")
		return
	}
	/*	_, err = exec.Command("go", "build", output+".go").Output()

		if err != nil {
			fmt.Println("could not compile")
			return
		}
	*/
}

func clean(src string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune("<>+-.,[]", r) {
			return r
		}
		return -1
	}, src)
}

func scan(src string, start int, char rune) int {
	var n int
	for _, c := range src[start:] {
		if c != char {
			return n
		}
		n++
	}
	return n
}

func transpile(src string) string {
	src = clean(src)
	program := `package main
import "fmt"
func main() {
var (
mem [30000]byte
dp  int
)
`
	for i := 0; i < len(src); i++ {
		switch src[i] {
		case '>':
			n := scan(src, i, '>')
			program += "dp += " + strconv.Itoa(n)
			i += n - 1
		case '<':
			n := scan(src, i, '<')
			program += "dp -= " + strconv.Itoa(n)
			i += n - 1
		case '+':
			n := scan(src, i, '+')
			program += "mem[dp] += " + strconv.Itoa(n)
			i += n - 1
		case '-':
			n := scan(src, i, '-')
			program += "mem[dp] -= " + strconv.Itoa(n)
			i += n - 1
		case '.':
			program += "fmt.Printf(\"%c\", mem[dp])"
		case ',':
			panic("input not implemented yet")
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
