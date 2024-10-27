## App Introduction

> This is a very basic `Conference Ticket Booking` CLI App built with `Go` Language.

## What is Go Language

> Golang is a `compiled`, `statically typed` programming language that's designed to be efficient, fast, and easy to learn.

<details>
<summary>Use cases</summary>

> Golang is used for a variety of applications, including web development, game development, cloud-based programming, data science, and creating command-line tools.

</details>

<details>
<summary>Features</summary>

> Golang is known for its clear documentation, versatility, and scalability. It's also usually faster than other high-level programming languages because its types don't change.

</details>

<details>
<summary>History</summary>

> Golang was first introduced to the public in 2009 and became open source in 2012. It was originally built to replace high-performance server-side languages like Java and C++.

</details>

<details>
<summary>Popularity</summary>

> Golang is popular with many tech giants, including `Google`, `Netflix`, `Twitch`, `Ethereum`, `Dropbox`, `Kubernetes`, `Docker`, and `Heroku`.

</details>

## GoLang Installation

Go to <u>[Official Website](https://go.dev/doc/install)</u> to download `go` as per your OS.

After Installation check in your `terminal`/`cmd`.

```bash
> go version
go version go1.23.2 windows/amd64
```

## How to setup first go project

To setup first go project, create a foler and inside it create a file called `main.go`.
After that run a command to initialize project or "module"

```bash
go mod init project_name
```

> 🗒️**Note:** This command will initialize project and create a `go.mod` file.

```mod
<!-- Content Inside go.mod -->

module learning01

go 1.23.2
```

## Packages in GoLang

> Packages are the most powerful part of the Go language. The <u>**_purpose of a package is to design and maintain a large number of programs by grouping related features together into single units_**</u> so that they can be easy to maintain and understand and independent of the other package programs.

> A Package is a collection of Go source files that reside in the same directory and have the same package declaration at the top.

```go
// Import one package
import "fmt"

// Import multiple packages
import (
    "fmt"
    "strings"
)
```

## Variables

```go
// Two ways to declare variable in golang
// Syntax:
// var variable_name data_type = value
// variable_name := value -> we cannot explicitly define type in it.

// Examples
adminName := "Niket Kumar Mishra"

var name string = "Niket";
age := 22;

const planetName string = "Earth"
```

## Array

```go
var bookings [50]string

// Initialized an empty array
var bookings = [50]string{}
```

## Slices

> `slice` in go, it is dynamic sized and using `array` under the hood.  
> assign emply slice
>
> ```go
> bookings := []string{}
> ```
>
> We <u>can not declare package level variable with this syntax</u>, we need to use <u>var/const</u>

```go
var bookings []string
```

> Create a slice with "make"

```go
var bookings make([]string)
```

## Maps

Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.

#### How to declare `map` in go?

type: `map[Key_Type]Value_Type`

> An empty map  
> example: `map[string]string{}`

> Map with key-value pair  
> map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

#### create a variable of `map` data type

```go
var user := map[string]string
```

#### create a `slice of map`

```go
var users := []map[string]string
```

```go
// Checking if the map is nil or not
var map_1 map[int]int

if map_1 == nil {
    fmt.Println("True")
} else {
    fmt.Println("False")
}
```

```go
// Creating and initializing a map
// Using shorthand declaration and
// using map literals
map_2 := map[int]string{
    90: "Dog",
    91: "Cat",
    92: "Cow",
    93: "Bird",
    94: "Rabbit",
}

fmt.Println("Map-2: ", map_2)
```

#### create a variable of map data type with `make()` funtion

```go
var users := make([]map[string])
```

#### create a slice of map

```go
var users := make([]map[string])
```

<details>
<summary>Features</summary>
- It is a reference to a hash table.
- Due to its reference type it is inexpensive to pass, for example, for a 64-bit machine it takes 8 bytes and for a 32-bit machine, it takes 4 bytes.
- In the maps, a key must be unique and always in the type which is comparable using == operator or the type which support != operator. So, most of the built-in type can be used as a key like an int, float64, rune, string, comparable array and structure, pointer, etc. The data types like slice and noncomparable arrays and structs or the custom data types which are not comparable don’t use as a map key.
- In maps, the values are not unique like keys and can be of any type like int, float64, rune, string, pointer, reference type, map type, etc.
- The type of keys and type of values must be of the same type, different types of keys and values in the same maps are not allowed. But the type of key and the type values can differ.
- The map is also known as a hash map, hash table, unordered map, dictionary, or associative array.
- In maps, you can only add value when the map is initialized if you try to add value in the uninitialized map, then the compiler will throw an error.
</details>

## Struct

A `structure` or `struct` in Golang is a user-defined type that **_allows to group/combine items of possibly different types into a single type_**. Any real-world entity which has some set of properties/fields can be represented as a struct. This concept is generally compared with the classes in object-oriented programming.

```go
// We have defined a custom data type
type UserData struct {
	firstName   string
	lastName    string
	emailId     string
	noOfTickets uint
    isStudent   bool
}
```

## Functions

> Functions are the heart of any programming language. That helps to avoid code repeatation.
>
> <u>**_Function is a reusable block of code that performs a specific task_**</u>

```go
// Function syntax in golang

func function_name(name string, age uint) {
    // Write login
}

// Example of a function to greet user
func greetUser(userName string){
    fmt.Printf("Welcome %v, We are very pleased to have you here!")
}
```

> 💡In GoLang `Functions` can return more than one value.

```go
func userDetails() {
    name := "Niket";
    age := 22;

    return name, age;
}

// This is how we can access variables returned from function
func main(){
    name, age := userDetails()
}
```

## Pointers in Go

Pointers in Golang is <u>**_a variable that is used to store the memory address of another variable_**</u>

```go
var x int = 100
var y *int = &x

// x -> 0x0201 -> 100
// y -> 0x0206 -> 0x0201 -> 100
// *y -> 0x0201 -> 100
// x == *y -> true
// &x == y -> true
```

![alt text](/assets/image.png)

## Package Level Variable

> Define outside function body and access throuout the package

> **Best Practice:** _Define Variable as "Local" as Possible, create the variable where you need._

```go
package main

import (
    "fmt"
    "strings"
)

// Package Level Variable
var name string = "Niket"

func main(){
    // This is local variable
    var age unit = 22

    fmt.Printf("My name is %v and my I am %v years old.", name, age)
}
```

## If/else conditions

<u>Use the `if` statement to specify a block of Go code to be executed if a condition is true.</u>

```go
if condition {
  // code to be executed if condition is true
}
```

<u>Use the `else` statement to specify a block of code to be executed if the condition is false.</u>

```go
if condition {
  // code to be executed if condition is true
} else {
  // code to be executed if condition is false
}
```

#### Nested Conditions

```go
if condition1 {
   // code to be executed if condition1 is true
} else if condition2 {
   // code to be executed if condition1 is false and condition2 is true
} else {
   // code to be executed if condition1 and condition2 are both false
}
```

## Loops

<u>Go language contains only a single loop that is for-loop</u>. **_A for loop is a repetition control structure that allows us to write a loop that is executed a specific number of times_**. In Go language, this for loop can be used in the different forms and the forms are:

Syntax:

```go
for initialization; condition; post{
       // statements....
}

// Examples
for i := 0; i < 4; i++{
    fmt.Printf("Author of this docs is Niket Kumar Mishra\n")
}

// Infinite for loop
for{
    fmt.Printf("Niket Mishra\n")
}

// As while loop
for condition{
    // statement..
}

for i < 3 {
    fmt.Printf("Value of i is %v \n", i)
    i += 2
}

// Simple range in for loop
 // Here `rvariable` is an array
for i, j:= range rvariable{
   // statement..
}

// Using for loop for strings
for index, chr:= range str{
     // Statement..
}

// For Maps
mmap := map[int]string{
    22:"Niket",
    33:"Developer",
}
for key, value := range mmap {
     fmt.Println(key, value)
}

// in go, "_" are used to identify unused variables, here "_" -> index
for _, booking := range bookings {
	var names = strings.Fields(booking)
	 = append(firstNames, booking.firstName)
}
```

> Read more about for loop here 👉 [For Loop in Go](https://www.geeksforgeeks.org/loops-in-go-language/)

## break and continue

<u>The break statement terminates the loop when it is encountered</u>. For example,

```go
for initialization; condition; update {
    // terminates the loop when initialization is equal to 2
    if(initialization == 2){
        break
    }
    fmt.Println(initialization)
}
```

In Go, <u>the continue statement skips the current iteration of the loop. It passes the control flow of the program to the next iteration</u>. For example,

```go
for initialization; condition; update {
    if condition {
        continue
    }
}
```

## Switch-case

```go
// syntax
switch optstatement; optexpression {
    case "expression1":
        // Statement
    case "expression2":
        // Statement
    default:
        // Statement
}
```

```go

// One Way
day := 2
switch day {
    case 1:
        fmt.Println("This is day one")
    case 2:
        fmt.Println("This is day two")
    default:
        fmt.Println("This is a random day")
}

// Second Way
switch city:="London"; city{
    case "New York":
        // Logic for new york bookings
    case "Singapore":
        // Logic for singapore bookings
    case "London":
        // Logic for london conference tickets

    // If booking logic for berlin and maxico is same
    case "Berlin", "Maxico City":
        // Logic for berlin or maxico city bookings
    default:
        fmt.Println("No valid city selected.")
}
```

## Import and Export Variables and Functions in Go

- Export a function or variable from a package so that it will be available for all packages
- Make `capitalize` first letter

## Scope of packages

- Global Scope
- Package Scope
- Local Scope

## Scope rules in GoLang

`Scope` is the **_region of program, where a defined variable and function can be accessed_**.

### Local

- Declaration **within function**:
  > Can be used only withing that function
- Declaration **within block**  
   (e.g. for if-else)
  > can be used only within that block

### Package

- **Define outside all functions**
  > Can be used everywhere in the same package

### Global

- Declaration outside all functions & uppercase first letter
  > Can be used everywhere across all packages
- Declaration **within block**  
   (e.g. for if-else)
  > can be used only within that block

## Concurency in `go`

### What Is Concurrency?

“<u>Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.</u>” — Rob Pike

> A `goroutine` is a function that is capable of running `concurrently` with other functions. To create a goroutine we use the keyword go followed by a function invocation:

```go
package main

import "fmt"

func f(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
  }
}

func main() {
  go f(0)
  var input string
  fmt.Scanln(&input)
}
```

**This program consists of two goroutines**. <u>The first goroutine is implicit and is the main function</u> itself. <u>The second goroutine is created when we call go f(0)</u>. _Normally when we invoke a function our program will execute all the statements in a function and then return to the next line following the invocation. With a goroutine we return immediately to the next line and don't wait for the function to complete. This is why the call to the `Scanln` function has been included; without it the program would exit before being given the opportunity to print all the numbers._

Goroutines are lightweight and we can easily create thousands of them. We can modify our program to run 10 goroutines by doing this:

```go
func main() {
  for i := 0; i < 10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}
```

> Read more about concurrency in `go`: [🔗Link](https://www.golang-book.com/books/intro/10)

## Goroutine

<u>_A goroutine is an independent function that executes simultaneously in some separate lightweight threads managed by Go_</u>. GoLang provides it to support concurrency in Go.

- Go is using, what's called a `green thread`
- **Abstraction** of an actual thread
- Managed by the go runtime, we are only interacting with these high level goroutines.
- Cheaper & lightweight
- You can run hundreds of thousands or millions goroutines without affecting the performance.
- Built-in functionality for go routines to talk with one another.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Here helloworld run in different thread which will not block the code execution
    go helloworld()
    time.Sleep(1 * time.Second)
    goodbye()
}

func helloworld() {
    time.Sleep(5 * time.Second)
    fmt.Println("Hello World!")
}

func goodbye() {
    fmt.Println("Good Bye!")
}
```

---

> Congratulation! 🔥🔥🔥 If you have read the entire docs carefully then you are good to start with go programming language.

If you like this docs then follow me on these platform to get more contents on `Web Development` and `Coding`.

## WaitGroups

To wait for multiple goroutines to finish, we can use a wait group.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Crate a waitGroup
var wg sync.WaitGroup

// This function will take 5 seconds to complete
func sendEmail() {
    time.Sleep(5 * time.Second)

    // reduce the waitlist once function has complete its execution
    defer wg.Done()
}

func main() {
    for i := 1; i <= 5; i++ {
        // Increment the waitlist
        wg.Add(1)
        // goroutine
        go sendEmail()
    }

    // wait for all the waitlist to be empty
    wg.Wait()
}
```

```
create a waitlist -> increment list -> reduce list -> wait for empty
```

### Connect with me

---

- [Github](https://github.com/niketvns)
- [LinkedIn](https://www.linkedin.com/in/niket-kumar-mishra-37ab5a215/)
