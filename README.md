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

Every Go program is made up of package. Programs start running in package `main`. 

Package are imported with the `import` keyword, in our example the `fmt` package is imported and used to print to the console.

    import "fmt"
    import "net/http"

Multiple imports can be grouped by a _factored_ import:

    import(
        "fmt"
        "net/http"
    )

After an package is imported the public types become available to the context that imports them and can be accessed by via the package name. E.q., `fmt.Println` or `http.Get("http://google.com")`.

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
    s := <- c          // receive from channel, and assign value to `s`

Channels are blocking by default, which means reading from a channel blocks until there is a value available, and sending to a channel is blocked until the value is received by another routine.

### Exersise

Instead of using `time.Sleep` to wait while rob and christian talk, use an channel to block the execution of the `main` function for at least `10` seconds.

Hints:

* The `time.After` method returns a channel imidiatly and sends the current time on the returned channel after the specified duration.
