# The COMP2310 Unofficial GOLANG Workshop

## Welcome!

Welcome! This workshop is a brief introduction to Go, intended at students studying COMP2310. It expects the reader has gone through the [Tour of Go](https://tour.golang.org/).

Go (sometimes referred to as Golang) is a concurrent programming language designed for simple, efficient, and safe software.

Go uses C-style syntax, similar to many languages you may be familiar with.

The key concurrent construct in Go is the *Goroutine*. These are concurrent threads that we can create with the `Go` keyword.

Let's imagine we have a function to calculate the maximum of a list and print the result:

```go
print_max(list)
```

If we want to run this statement concurrently, we can run:

```go
go print_max(list)
```

This will create a new thread to run this function. In the meanwhile, our main thread will continue running.

For example, imagine the following code:

```go
f()
go g()
h()
```

This will trigger the following execution process:

```
f()
|
|
go g()
|     \
|      |
h()   g()
```

Importantly, there are no guarantees as to *when* `g` is run. This is important to consider when writing code. For example:

```go
func test() {
	time.Sleep(10 * time.Second)
	fmt.Println("Test!")
}

func main() {
	go test()
}
```
In this case, running this will *not* output anything. This is because the `main` function will terminate immediately, which in turn kills the child process.

## Our first Go program.

Let's begin with our first program.

Create a folder called `hello_world` and create a file within it called `hello_world.go`. Start by adding `package main` to the beginning of the file. This indicates that the program is executable rather than a library. Note that Go assumes all files in the same directory belong to the same package, hence why we create a new folder for each exercise.

We next need to create the main function for our program. Create a function named `main` and add:
```go
fmt.Println("Hello world!")
```
inside it. If you're using an IDE, it should automatically import the necessary package. Otherwise, add
```go
import (
	"fmt"
)
```
after the `package main`.

Finally, we need to run this! `cd` into `hello_world` and the file:
```shell
$ go run hello_world.go
Hello world!
```

## Harder problems.

From here, we'll swap to live teaching. This repo contains a number of practice exercises of varying difficulty and topic. Feel free to choose one that looks interesting, and give it a shot!

In approximately increasing level of difficulty:
* [`waiting`](waiting/waiting.go) [!] - use WaitGroups to let the main process wait for a child thread to complete.
*  [`token_ring`](token_ring/token_ring.go) [!!!] - finish an implementation of a token ring. This has an easy solution and a high efficiency solution.

&copy;Robert Jeffrey, 2021