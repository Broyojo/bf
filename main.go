package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	var name string

	flag.StringVar(&name, "name", "", "The name of the file")
	flag.Parse()

	src, err := ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	program := compile(src)
	newName := strings.Replace(name, ".b", ".go", 1)
	if err := WriteFile(newName, program); err != nil {
		fmt.Println("compilation failed")
		return
	}

	_, err = exec.Command("go", "build", newName).Output()

	if err != nil {
		fmt.Println("compilation failed")
		return
	}

	_, err = exec.Command("rm", "-rf", newName).Output()

	if err != nil {
		fmt.Println("could not delete temporary file")
		return
	}
}

func compile(src string) string {
	program := `package main
import "fmt"
func main() {
var mem [30000]byte
dp := 0
`
	for _, instr := range src {
		switch instr {
		case '>':
			program += "dp++"
		case '<':
			program += "dp--"
		case '+':
			program += "mem[dp]++"
		case '-':
			program += "mem[dp]--"
		case '.':
			program += "fmt.Printf(\"%c\", mem[dp])"
		case ',':
			panic("input reading not implemented yet")
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
