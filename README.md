# Learn a new language Go

## Installation

We assume that you have installed Go as described in <http://golang.org/doc/install>.

Install the package that matches you platform an set the following environment variables:

* Env var `GOROOT`, points to the Go binaries, e.q.: `/usr/local/go`
* Env var `GOROOT`, should be added to your `PATH`
* Evn var `GOPATH`, should be set to a directory that will hold your Go workspace like `$HOME/go` and should contains 3 folders: `src`, `pkg`, `bin`.
* Now `go version` should output something like: `go version go1.3 linux/amd64`.

## Clone de workshop repository

    $ go get github.com/pjvds/goworkshop

Code will be cloned into your Go workspace: `$GOPATH/src/github.com/pjvds/goworkshop`.

1. `cd` into the `goworkshop`
2. open the `readme`
3. each exercise has its own directy
4. code can be run by `go run main.go` from the exercise directories

## 01 Hello World

Every Go program is made up of packages. Programs start running in package `main`. 

Package are imported with the `import` keyword, in our example the `fmt` package is imported and used to print to the console.

    import "fmt"
    import "net/http"

Multiple imports can be grouped by a _factored_ import:

    import(
        "fmt"
        "net/http"
    )

After a package has been imported the public types become available to the context that imports them and can be accessed via the package name, the last part of the import: e.g., `fmt.Println` or `http.Get("http://google.com")`.

### Exercise

Change the output of the program from `hello world` to `my favorite number is 3`, where `3` is a random number from `1` to `10`.

Hints:

1. The `rand` package is located at `math/rand`
2. The `Intn` method from the `rand` package can be used to get an positive random `int`.

## 02 Functions

Functions in Go are defined with the `func` keyword followed by the name of the function, zero or more arguments and finally the return values, e.q.:

    func name(a int, b int) (string, bool) {
        return "foobar", true
    }

### Exercise

Add the required function to make the code compile and print "hello world" to the console.

## 03 Odd or even

### Variables

Go has different ways to declare and initialize variables:

Declare and assign:

    var a int
    a = 1

Declare and assign one liner:

    var b int = 2

Short declaration and assignment:
    c := 3

### If

The `if` statement looks as it does in C or java, except that the ( ) are gone and the { } are required.

    if 5 < 10 {
        fmt.Printf("5 is smaller than 10")
    }

In contrast to C or java, Go does allow you to start an `if` statement with a short statement to execute before the condition. Variables declared by the statement are only in scope intul the end of the `if`.

    if n := getNumber(); n < 10 {
        fmt.Printf("%v is smaller than 10", n)
    }

### Exercise

Add the required function to make the code compile and print make sure it returns the correct result for 13, or any other int value.

## 04 Go routines

A go routine is a lightweight thread managed by the Go runtime. You can invoke a function in a new go routine with the `go` keyword:

    go talk("rob")

### Exercise

Let rob and christian talk at the same time by executing the `talk` function for both in go routines. Let them talk for at least `10` seconds.

Hints:

* All go routines get terminated when `main` returns.
* `time.Sleep` pauses the current routine for at least the duration.

## 05 Channels

Channels are a typed conduit through which you can send and receive values.

You create them with the `make` keyword:

    // Create a channel for string values.
    c := make(chan string)

You can send and receive with the `<-` operator:

    c <- "hello world" // Send string to channel `c`
    v := <- c          // receive from channel, and assign value to `v`

### Exercise

Change the code so that Rob and Christian do not talk to directly to the console (`fmt.Printf`), but to a channel instead.

* Create a new channel in the `main` method.
* Pass the channel as an argument to the `talk` method.
* Change the `talk` method body so it talks into the channel instead of to the console directly.
* Read from the channel in an endles loop from the `main` method and print the values to the console with `fmt.Printf`.

Hints:

* Use `fmt.Sprintf` instead of `fmt.Printf` to get a string value instead of printing it to the console.
* You can create a endless loop with the `for` keyword: `for{ ... }`.
* Don't worry about the endles loop you create, `ctrl`+`c` should kill your process.

## 07 Select

The select statement lets a go routine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random when multiple are ready. 

    select {
        case v := <-chanA:
            fmt.Printf("value from chanA: %v", v)
        case v := <-chanB:
            fmt.Printf("value from chanB: %v", v)
    }

### Exercise

Rob and Christian's talk is not that interesting, definitely not interesting enough to listen to it endlessly.

* Create a channel with `time.After` that creates a channel and waits in a go routine for the duration to elapse and then sends the current time on the returned channel.
* In the `for` loop use `select` to select from both channels.
* In case of a value from `c`, print to the console.
* In case of a value from the timeout channel, print "bye!" to the console and exit the program.

Hints:

* A program exits when the `main` method returns.

## 08 Chinese whisper

How expensive are go routines and channels? Lets find out by doing a chinese whisper game. The initial exercise code provides two methods, `start` and `whisper`. The `whisper` method gets the value from the `from` channel and, adds `1` to it, and writes it to the `to` channel.

### Exercise

* Execute the code to see it prints `2`.
* Change the `main` method so that is creates `10000` go routines and whispers between them (like the circle in the classroom from your childhood).
* Time the execution via the commandline: `time go run main.go` (first time you do this it also includes build time, so run twice).

## 09 Struct

A struct is a collection of fields:

    type Foo struct {
        Bar int
        Baz int
    }

Structs can be initialized the following:

    foo := Foo{
        Bar: 1,
        Baz: 2, // Yes, trailing comma
    }

### Exercise

* Define a struct called `Rect` with two `int` fields; `Width` and `Height`.
* Initialize a new instance variable of the `Rect` type called `r` and set the fields to a value of your liking.
* Calculate the surface area of the rectangle and print the result to the console.

## 10 Methods

Functions in Go can be attached to an type to make them a method.

    type Foo struct {
        Bar int
        Baz int
    }

    func (f *Foo) GetHighest() int {
        if f.Bar > f.Baz {
            return f.Bar
        }
        return f.Baz
    }

### Exercise

* Add a method to the `Rect` type that calculates the area and returns the result.
* Use this method to calculate the area and print the value from the `main` function.

## 11 Interface

Go is well known of its flexible type system. One good example of this is the way Go handles interfaces. Here is how an interface is defined:

    type Highester interface {
        GetHighest() int
    }

When a type has all methods defined by an interface, then it implements that interface automatically. 
So to implement the `Highester` interface you just attach a method that matches the signature to your type:

    type Foo struct {
        Bar int
    }

    func (f *Foo) GetHighest() int {
        if f.Bar > f.Baz {
            return f.Bar
        }
        return f.Baz

    }

### Exercise

* Implement the `Area` interface for the `Rect`, `Square` and `Circle` types.

## 12 Errors

Go handles errors differently than most other modern languages. Go does not think that errors are an exceptional case. It is something you just need to deal with. So it uses the power of multiple return values, and the strictness that you cannot assign a variable and not use it to encourage defensive programming.

In this example the `GetNumber` method returns a number and an error. The convention is that the normal results go first, and the last result is an error object. If the error is not nil, you know something went wrong.

### Exercise

* Try to build the code without changing it.
* Try to assign the error result to a variable as well and retry to build it.
* Handle the error value if it is not nil, by printing it to the console.

## Concurrent Prime Generator

It's time to take all you learned in practise and try to complete this single goals:

* Creating a prime generator that can create all of the primes below 2 million, in less than a minute.

Hints:

Set the environment variable `GOMAXPROCS` to the number of cores in your machine to execute the go routines in parallel.
