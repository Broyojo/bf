package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("You must specify a Brainf*ck source code file")
	}

	source, err := ReadSourceCode(args[0])
	if err != nil {
		log.Fatalln("Could not read source code:", err)
	}

	transpiled, err := Transpile(source)
	if err != nil {
		log.Fatalln("Could not transpile source code:", err)
	}

	err = WriteSourceCode(transpiled, "program.go")
	if err != nil {
		log.Fatalln("Could not write source code:", err)
	}
}

func ReadSourceCode(path string) (string, error) {
	source, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(source), nil
}

func Transpile(source string) (string, error) {
	transpiled := `package main

import (
	"fmt"
	"os"
)

func main() {
	mem := make([]byte, 30000)
	var dp int
`
	for _, c := range source {
		switch c {
		case '>':
			transpiled += "dp++"
		case '<':
			transpiled += "dp--"
		case '+':
			transpiled += "mem[dp]++"
		case '-':
			transpiled += "mem[dp]--"
		case '.':
			transpiled += "fmt.Printf(\"%c\", mem[dp])"
		case ',':
			transpiled += "read(mem, dp)"
		case '[':
			transpiled += "for mem[dp] != 0 {"
		case ']':
			transpiled += "}"
		default:
			continue
		}
		transpiled += "\n"
	}

	transpiled += "}\n"

	transpiled += `func read(mem []byte, dp int) {
	b := make([]byte, 1)
	if _, err := os.Stdin.Read(b); err != nil {
		panic(err)
	}
	mem[dp] = b[0]
}`

	return transpiled, nil
}

func WriteSourceCode(transpiled, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write([]byte(transpiled)); err != nil {
		return err
	}

	return nil
}
