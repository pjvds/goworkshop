# Learn a new language Go

## Installatie

In deze workshop gaan we er vanuit dat je Go beschikbaar is op je machine en de volgende environment variable zijn geset:

* Env var `GOROOT`, verwijst naar je Go binaries, bijvoorbeeld: `/usr/local/go`
* Env var `GOROOT`, is aan je `PATH` toegevoegd.
* Evn var `GOPATH`, verwijst naar `$HOME/go` en bevat de volgende drie mappen: `src`, `pkg`, `bin`.
* Het commando `go version` geeft het volgende terug: `go version go1.3 linux/amd64`.

Voor meer informatie over het installeren van Go, zie: <http://golang.org/doc/install>.

## Clone de workshop repository

Voer het volgende command uit om de workshop repository via het `go` command op te halen:

    $ go get github.com/pjvds/goworkshop

De code is gecloned in: `$GOPATH/src/github.com/pjvds/goworkshop`.

## Hello World

Every Go program is made up of packages. Programs start running in package `main`. 

Package are imported with the `import` keyword, in our example the `fmt` package is imported and used to print to the console.

    import "fmt"
    import "net/http"

Multiple imports can be grouped by a _factored_ import:

    import(
        "fmt"
        "net/http"
    )

After an package is imported the public types become available to the context that imports them and can be accessed by via the package name, the last part of the import: e.q., `fmt.Println` or `http.Get("http://google.com")`.

### Execsise

Change the output of the program from `hello world` to `my favorite number is 3`, where `3` is a random number from `1` to `10`.

Hints:

1. The `rand` package is located at `math/rand`
2. The `Intn` method from the `rand` package can be used to get an positive random `int`.

## Functions

Functions in Go are defined with the `func` keyword followed by the name of the function, zero or more arguments and finally the return values, e.q.:

    func name(a int, b int) (string, bool) {
        return "foobar", true
    }

### Execsise

Add the required function to make the code compile and print "hello world" to the console.

## Odd or even

### Variables

Go has different way to declare and initialize variables:

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

In contrast to C or java, does Go allow you to start an `if` statement with a short statement to execute before the conditation. Variables declared by the statement are only in scope intul the end of the `if`.

    if n := getNumber(); n < 10 {
        fmt.Printf("%v is smaller than 10", n)
    }

### Execsise

Add the required function to make the code compile and print make sure it returns the correct result for 13, or any other int value.

## Go routines

A goroutine is a lightweight thread managed by the Go runtime. You can invoke a function in a new go routines with the `go` keyword:

    go talk("rob")

### Execsise

Let rob and christian talk at the same time by executing the `talk` function for both in a go routine. Let them talk for at least `10` seconds.

Hints:

* All go routines get terminated when `main` returns.
* `time.Sleep` pauses the current routine for at least the duration.

## Channels

Channels are a typed conduit through which you can send and receive values.

You create them with the `make` keyword:

    // Create a channel for string values.
    c := make(chan string)

You can send and receive with the `<-` operator:

    c <- "hello world" // Send string to channel `c`
    v := <- c          // receive from channel, and assign value to `v`

### Exersise

Change the code so that Rob and Christian do not talk to directly to the console (`fmt.Printf`), but to a channel instead.

* Create a new channel in the `main` method.
* Pass the channel as an argument to the `talk` method.
* Change the `talk` method body so it talks into the channel instead of to the console directly.
* Read from the channel in an endles loop from the `main` method and print the values to the console with `fmt.Printf`.

Hints:

* Use `fmt.Sprintf` instead of `fmt.Printf` to get a string value instead of printing it to the console.
* You can create a endles loop with the `for` keyword: `for{ ... }`.
* Don't worry about the endles loop you create, `ctrl`+`c` should kill your process.

## Select

The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. 

    select {
        case v := <-chanA:
            fmt.Printf("value from chanA: %v", v)
        case v := <-chanB:
            fmt.Printf("value from chanB: %v", v)
    }

### Exersise

Rob and Christian talk is not that interesting, definitely not interesting enough to listen to it endlessly.

* Create a channel with `time.After` that creates a channel and waits in a goroutine for the duration to elapse and then sends the current time on the returned channel.
* In the `for` loop use `select` to select from both channels.
* In case of a value from `c`, print to the console.
* In case of a value from the timeout channel, print "bye!" to the console and exit the program.

Hints:

* A program exists when the `main` method returns.

## Chinese whisper

How expensive are goroutines and channels? Lets find out by doing a chinese whisper game. The initial exersise code provides two methods, `start` and `whisper`. The `whisper` method gets the value from the `from` channel and, adds `1` to it, and writes it to the `to` channel.

### Exersise

* Execute the code to see it prints `2`.
* Change the `main` method so that is creates `10000` goroutines and whispers between them (like the circle in the classroom from your childhood).
* Time the execution via the commandline: `time go run main.go` (first time you do this it also includes build time, so run twice).

## Struct

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

### Exersise

* Define a struct called `Rect` with two `int` fields; `Width` and `Height`.
* Initialize a new instance variable of the `Rect` type called `r` and set the fields to a value of your liking.
* Calculate the surface area of the rectangle and print the result to the console.

## Methods

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

### Exersise

* Add a method to the `Rect` type that calculates the area and returns the result.
* Use this method to calculate the area and print the value from the `main` function.

## Interface

Go is well known of its flexible type system. One good example of this is the way Go handles interfaces. Here is how an interface is defined:

    type Highester interface {
        GetHighest() int
    }

When a type has all methods defined by an interface, then it implements that interface automaticly. 
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

### Exersise

* Implement the `Area` interface for the `Rect`, `Square` and `Circle` types.

## Errors

Go handles errors differently that most other modern languages. Go does not think that errors are an exceptional case. It is something you just need to deal with. So uses the power of multiple return values, and the strictness that you cannot assign a variable and not use it to encourage defensive programming.

In this example the `GetNumber` method returns a number and an error. The convention is that the normal results go first, and the last result is an error object. If the error is not nil, you know something went wrong.

### Exersise

* Try to build the code without changing it.
* Try to assign the error result to a variable as well and retry to build it.
* Handle the error value if it is not nil, by printing it to the console.
