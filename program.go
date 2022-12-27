package main

import (
	"fmt"
	"os"
)

func main() {
	mem := make([]byte, 30000)
	var dp int
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
for mem[dp] != 0 {
dp++
mem[dp]++
dp++
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
dp--
dp--
mem[dp]--
}
dp++
mem[dp]++
mem[dp]++
dp++
dp++
mem[dp]++
dp--
for mem[dp] != 0 {
mem[dp]--
for mem[dp] != 0 {
dp++
dp++
mem[dp]++
dp--
dp--
mem[dp]--
}
mem[dp]++
dp++
dp++
}
dp++
mem[dp]++
for mem[dp] != 0 {
mem[dp]--
dp--
dp--
dp--
for mem[dp] != 0 {
mem[dp]--
dp++
for mem[dp] != 0 {
mem[dp]++
for mem[dp] != 0 {
mem[dp]--
}
mem[dp]++
dp++
mem[dp]++
mem[dp]++
dp++
dp++
dp++
mem[dp]--
dp--
dp--
}
dp--
for mem[dp] != 0 {
dp--
}
dp++
dp++
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
for mem[dp] != 0 {
dp--
dp--
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
mem[dp]++
dp++
dp++
mem[dp]--
}
mem[dp]++
dp--
dp--
mem[dp]++
mem[dp]++
fmt.Printf("%c", mem[dp])
for mem[dp] != 0 {
mem[dp]--
}
dp--
dp--
}
dp++
fmt.Printf("%c", mem[dp])
dp++
mem[dp]++
for mem[dp] != 0 {
dp++
dp++
}
dp++
mem[dp]++
}
}
func read(mem []byte, dp int) {
	b := make([]byte, 1)
	if _, err := os.Stdin.Read(b); err != nil {
		panic(err)
	}
	mem[dp] = b[0]
}