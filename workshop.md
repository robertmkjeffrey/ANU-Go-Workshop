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

**Checkpoint 1: run `hello_world.go`.**

## Our first *concurrent* Go program.

Alright, everything so far looks normal. Let's start writing some concurrency.

We'll tackle a problem inspired by [COMP2310 Lab 3](https://cs.anu.edu.au/courses/comp2310/labs/03-tasks/). We'll create multiple tasks with some adding and some subtracting from a shared variable.

One important difference between Ada and Go is the reduced focus on shared variables. In Ada, a task is able to access and modify variables defined in its scope.

```ada
procedure Main is
    Sum : Natural := 50;

	task Counter;
	task body Counter is
	begin
		-- This changes the value outside the task.
	    Sum = Sum + 1;
	end Counter;
begin
	null;
end Main;
```

This is not generally true in Go! As is often quotes, Go follows the philosophy:

> [Do not communicate by sharing memory; instead, share memory by communicating.](https://golang.org/doc/effective_go)

So what does this mean? Unlike Ada, Go is not designed with powerful data protection semantics (e.g. protected objects) at the core. Instead, the fundamental tool we have is a channel.

Channels can act as one of two things: a synchronous messaging passing point or a concurrent queue. The simplest definition creates the synchronous form. 
```go
make(chan int) // Synchronous
make(chan int, 0) // Synchronous
make(chan int, 5) // Buffered with 5 elements
```

In either case, we have two operations for a channel. We can either add insert and element or retrieve an element. These are written as follow:

```go
c := make(chan int)
c <- 0 // Put an element in C
x := <- c // Read an element from C
```

Note that this will deadlock on a sequential process - the channel is synchronous, so any write is blocked until a matching read and vice versa.

Now, let's implement some shared memory through communicating. For this problem, we'll have multiple tasks each aiming for a different target number. Each will read the current value, check if it's larger or smaller than its target, and then add or subtract 1 respectively. 

Open [`counters`](counters/counters.go) and read through the code. Each "counter" is implemented by the `counter` function running on a new goroutine. These will write the requested changes into the `update_chan`. As these channels are unbuffered, the `counter` must wait until the request is handled before it can request additional updates.

The updates to the shared value `count` are handled by the `count_manager` function running on its own goroutine. This accepts requests on `update_chan` and updates the variable accordingly. It additionally has the ability to shut down if nobody requests changes for a certain amount of time.

If you run this code, you'll see that it manages to update the value as desired! Unfortunately, the `counter`s can't see the changes to the `count` so they continue forever. Your task is to allow the `counter`s to view the updated `count` values using channel-based communication.

**Checkpoint 2: fix the `counters` program.**

## Using shared data.

You might be wondering why we have to do so much work to get a single shared variable working - why can't we just use a shared variable? 

The benefit of communication-based memory sharing is that we prevent the classic issues caused by concurrent accesses to a piece of data. Architecting a system to focus on interacting tasks, rather than shared data, can provide a lot of safety and help make hard code-bases safe!

That said, this case is one where shared data would be reasonable. The only requirement is that we protect it successfully. To do this, we can use the `sync` package from the standard library. In this case, we're particularly interested in the `sync.RWLock` object. This provides three main functions: `Lock` will block all other tasks from gaining a lock on the object, and `Unlock` will release any current locks. The `RLock` provides the ability for a read-only lock; it will block any `Lock` calls, but allows other `RLock` calls to succeed.

Let's create a protected piece of shared data using this. We could simply use them separately:
```go
var lock sync.RWMutex
var n int = 0

lock.Lock()
n += 1
lock.Unlock()
```
but this would be confusing once we have multiple variables to protect. Instead, let's embed the lock in a struct:

```go
type ProtectedInt struct {
	sync.RWMutex
	n int
}
```

which allows us to call the lock functions on the struct itself. This gives us the following usage:

```go
var count ProtectedInt

count.Lock()
count.n += 1
count.Unlock()
```
which ensures we always use the correct lock!

Your task now is simple: rewrite your solution from `counters` to use this shared variable. Place this in the `counters_v2` folder and name the file `counters_v2.go`.

There are two issues you may run into. First, how do you share variables between tasks? This will require the use of pointers. Here's an example:

```go
// The "*" indicates this will be a pointer to an int.
func addOne(n *int) {
	// "*" dereferences the pointer to get the value.
	*n = *n + 1
	return
}

func main() {
	n := 0
	// "&" creates a pointer to a variable
	addOne(&n)
	fmt.Println(n)
}
```

The other issue you may run into is deciding when to end the program. The best way to do this is to use a WaitGroup; [this resource gives examples of using one](https://gobyexample.com/waitgroups).

**Checkpoint 3: rewrite the `counters` program using protected shared variables.**

## Harder problems.

From here, we'll swap to live teaching. This repo contains a number of practice exercises of varying difficulty and topic. Feel free to choose one that looks interesting, and give it a shot!

In approximately increasing level of difficulty:
* [`waiting`](waiting/waiting.go) [!] - use WaitGroups to let the main process wait for a child thread to complete.
*  [`token_ring`](token_ring/token_ring.go) [!!!] - finish an implementation of a token ring. This has an easy solution and a high efficiency solution.

&copy;Robert Jeffrey, 2021