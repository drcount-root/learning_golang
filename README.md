# Process of setup & execution

## To create a go module

`go mod init example.com/hello`

## To install modules

`go mod tidy`

## To run

`go run .`

go run main.go

go build main.go

./main

---

## Notes for Packages

### Internal Packages

In go package name (fileops) & folder name (fileops) should be same, filename can be different

### To install external package

`go get github.com/Pallinder/go-randomdata`

---

## Go is not a dynamically typed language. It is a statically typed with type inference

Let’s frame it precisely, the way a system architect thinks:

1. Statically Typed

Go determines the type of every variable at compile time, not at runtime.
Even if you use:

x := 10

Go infers the type int, and after that, x is strictly an int.
You cannot later do:

x = "hello" // compile-time error

2. Type Inference

The := syntax lets the compiler infer the type so you don’t have to write it:

name := "Bro" // inferred as string
pi := 3.14 // inferred as float64

The key detail:
Type inference does not make Go dynamic; it only removes verbosity.

3. How it differs from dynamically typed languages

Languages like Python and JavaScript decide variable types at runtime.
You can do:

x = 10
x = "hello"

In Go, this is impossible. Types are fixed once assigned.

4. Go has some dynamic behavior via interfaces

If you use interface{}, Go can hold values of any type, but even then, the actual type is preserved internally.
This is controlled dynamic capability, not dynamic typing.

Final summary:

Go = statically typed + type inference + limited dynamic features via interfaces
