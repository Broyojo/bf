package main
import "fmt"
func main() {
var (
mem [30000]byte
dp  int
)
}
mem[dp] += 1
dp += 1
dp -= 1
}
fmt.Printf("%c", mem[dp])
mem[dp] -= 1
for mem[dp] != 0 {
fmt.Printf("%c", mem[dp])
dp += 1
fmt.Printf("%c", mem[dp])
}
fmt.Printf("%c", mem[dp])
mem[dp] -= 1
dp += 1
fmt.Printf("%c", mem[dp])
dp -= 1
}
dp -= 1
for mem[dp] != 0 {
}
}
}