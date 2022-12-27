# Brainf*ck to Golang Transpiler
![](diagram.jpeg)

## Description
This program can transpile [Brainf\*ck](https://en.wikipedia.org/wiki/Brainfuck), an esoteric programming language, to Golang. For each symbol in the Brainf\*ck code, I translate it to the corresponding Go code. I chose to do it this way because an interpreter would have been slower at executing the code. The transpiler will emit a `program.go` file which is the transpiled version of the input source code. 

## How to Run
```
$ go run main.go path/to/source
$ go run program.go
```
There are many example programs located in the [programs](https://github.com/Broyojo/bf/tree/master/programs) directory.


## Example
```
$ go run main.go programs/mandelbrot.b
$ go run program.go
```
![](example_output.png)