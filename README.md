# Resources

- https://github.com/dariubs/GoBooks
- Book: Go Programming Language, The (Addison-Wesley Professional Computing Series)
- Books: https://realtoughcandy.com/best-golang-books/

# Reasons GO was created

- Multi Threading (each thread is executing one task)
- Concurrency: dealing with lots of things at once
- Concurrency in GO is cheap and easy
- Go was designed to run on multiple cores and built to support concurrency

 

# Basics

## How to run

`cd <folder with main.go file>`

`go run main.go` 

## Commands

`go run` ---> compiles and runs one or more files (creates an executable and runs it)

`go build` ---> compiles multiple go source files (creates an executable but does not run it)

`go fmt` ---> formats all the code in each file in the current directory

`go install` ---> compiles and installs a package

`go get` ---> downloads the raw source code of someone else's package

`go test` ---> runs any tests associated with the current project

`go mod init <module path>` ---> creates new module

- initializes the go.mod file
- describes the module: **with name/module path** and **go version** used in the program
- if the module is booking-app then that is the name of the import path as well.

`go get -u go get -u <pkg url>` 

- Downloads package based on the url.
- for example `go get -u go get -u github.com/gorilla/mux` in order to then use package from the url
- The file go.sum will be updated

## Types of packages

- Executable: Generates a file that we can run, declare `package main` and must have a func named main. Defines a package that can be compiled and then executed.
- Reusable: Code used as helpers, reusable logic. Can be used as a dependency.

## Packages that come with go

https://golang.org/pkg/

## Basic Types

Basic Types (υπάρχουν κι άλλοι τύποι)

- bool
- string
- int
- float64

## Constant

Syntax : `const name = "Fotis"`

## Go is Static type language

Κάθε μεταβλητή έχει ένα συγκεκριμένο τύπο, αυτός πρέπει να δηλωθεί ρητά. Όπως σε Java.

Σε αντίθεση με Javascript που είναι Dynamic Type Lanuage όπου μπορώ σε ένα string μετά να βάλω number.

# Array and Slice

Whenever we create a slice, GO will automatically create two data structures:

1. An array
2. A structure that records the length of the slice, the capacity of the slice, and a reference to the underlying array.



## Syntax

- Create empty array of maps: `var bookings = make([]map[string]string, 10)` where 10 is the number of elements
- Create empty slice of maps: `var bookings = make(map[string]string)`



## **Array VS Slice**

- **Array**: 
  - Primitive data structure
  - Fixed length list of things
  - Rarely used directly
  - Syntax 
    - `var bookings = [50]string{"Fotis", "Nana", "Peter"}`  ---> creates array of 50 string elements and initialize some values
    - `var bookings [50]string` ---> creates array of 50 string elements empty
    - `len(bookings)` --> length of array
    - `bookings[0] = "Fotis"`---> set value of array
  
- **Slice**: 

  - Slice is an abstraction of array
  - Slices are more flexible and powerful - variable length or get a sub array of its own
  - Slices are index based and have a size, but is resized when needed
  - Used 99% of the time for lists of elements
  - Syntax
    - `var bookings []string`  or `bookings := []string{}` ---> create slice of strings
    - `len(bookings)` --> length of slice
    - `bookings = append(bookings, firstName+" "+lastName)` ---> set value of slice
  - Array that can grow or shrink. Every element  must be the same type.
  
    - Example


    package main
    
    import "fmt"
    
    func main() {
    	cards := []string{"Ace of Diamonds", newCard()}
    	cards = append(cards, "Six of Spades")
    
    	for i, card := range cards {
    		fmt.Println(i, card)
    	}
    }
    
    func newCard() string {
    	return "Five of Diamonds"
    }


## Parse Slice

**Example**: Parse Slice using index

```go
for i, card := range d {
		fmt.Println(i, card)
	}
```

**Example**: Parse Slice - and say that I don't want to use the index. It is using the underscore.

```go
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
```

**Example**

If cards is an array of 4 elements {0,1,2,3}

`cards[:3]` ---> contains elements 0,1,2

`cards[3:]` ---> contains element 3

`cards[0:2]` ---> contains elements 0,1 but does not contain the 2

# Struct

Data structure. Collection of properties that are related together.



Syntax

```
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

```



# Map

Key - value structure.

Used to map unique keys to values

All keys must have the same type.

All values must have the same type.

Keys and values do not have to have the same type.



- Declare empty map: `var userData = make(map[string]string)` which is `map[key_type]value_type`
- Save/Add to map: `userData["firstName"] = firstName`
- Get the value of an entry with key = firstName: `booking["firstName"]`



## Example: Declare empty map and add data

```
var userData = make(map[string]string)
userData["firstName"] = firstName
userData["lastName"] = lastName
userData["email"] = email
```





# Reminders

- We can initialize a variable outside of a function, we just can't assign a value to it.
- Files in the same package can freely call functions defined in other files, without using import.
- Default values of types
  - string: ""
  - int: 0
  - float: 0
  - bool: false

## Map vs Struct

### Package Level Variables

Can be declared only using var. So this is correct `var name= "Fotis"` but this is not correct `name := "Fotis"` for package level variables.

### Map

- All keys must be the same type
- Used to map unique keys to values
- Use to represent a collection of related properties
- Don't need to know all the keys at compile time
- Keys are indexed - we can iterate over them
- Reference type

### Struct

- Values can be of different type
- You need to know all the different fields at compile time
- Keys don't support indexing
- Use to represent a thing with a lot of different properties
- Value type





## String manipulation

- `strings.Fields()` --> splits the string with white space as separator
- `strings.Contains(email, "@")` ---> returns if email(type string) contains the character @
- `var ticket = fmt.Sprintf("%v tickets for %v %v.\n", userTickets, firstName, lastName)` ---> create formatted string and save it to a string variable with name `ticket`



## Syntax: Function Returns Multiple Variables

```
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= uint(remainingTickets)

	return isValidName, isValidEmail, isValidTicketNumber
}
```

and the code that will call this function

```
isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
```



## Organize code in Packages

### Scenario: Multiple files, same package

If we have files `main.go` and `helper.go` , where `main.go` calls function of `helper.go` and both files belong to the package `main` then in order to run the code, I must use the command `go run main.go helper.go`, or `go run .` which is asking to execute all files from the same folder.

- No export of function is needed
- No import of a package is needed

### Scenario: Multiple files, multiple packages

If we have files `main.go` and `helper/helper.go` , where `main.go` calls function of `helper/helper.go` . Also `main.go` belongs to package `main` and `helper/helper.go` belongs to package `helper`. 

Also if `go.mod` has `module booking-app`

- `helper/helper.go` must export the function by just capitalizing the first name of the function. So the declaration of the function will look like the following:  `func ValidateUserInput(....)`
- `main.go` must have import from `booking-app/helper`
- `main.go` can call the function using the syntax `helper.ValidateUserInput(....)`

### Export function, or variable

In order to export function or variable from one package to another, then just capitalize the first letter of the name.








# Conversion

## String ---> Slice of Bytes

- `[]byte("Hi there!")` ---> returns [72 105 32 116 104 101 114 101 33]

## Slice of String ---> String

Use function Join contained in the *strings* GO lang packages.

- `strings.Join([]string(d), ",")` ---> d is the slice of strings, and "," is the separator contained in the string result



### Uint ---> String

`strconv.FormatUint(uint64(userTickets), 10)`---> where userTickets is a uint

Takes uint value and formats it to a string as a decimal number. The second parameter `10` means that the result will be a decimal number. If I used `16` then the result would be a hexadecimal number.



# Pointer and Value

What are they doing:

- **&variable** Give me the memory address/pointer of the value this variable is pointing at 

- ***pointer** Give me the value this memory address is pointing at



In other words:

- **&value** : turn value into address/pointer
- ***address** : turn address/pointer into value



## Theory: Value Types vs Reference Types

**Value Types:** Use pointers to change these things in a function. In other words: whenever you pass those types into a function --> Go creates a <u>copy</u> of each argument, and these copies are used inside the function.

- int
- float
- string
- bool
- structs

**Reference Types:** Don't have to worry about pointers with these.

- slices
- maps
- channels
- pointers
- functions



## Example with Primitive

The following code will NOT update the myVar value.

```go
func main() {
	myVar := 0

	update(myVar)

	fmt.Println(myVar) // print 0
}

func update(i int) {
	i = 1
}
```

The following code will update the myVar value.

```go
func main() {
	myVar := 0

	update(&myVar)

	fmt.Println(myVar) // print 1
	fmt.Println(&myVar) // print memory address
	fmt.Println(*&myVar) // print 1
}

func update(i *int) {
	*i = 1
}
```



## Example with Struct

Call the receiver function (option 1)

```go
jimPointer := &jim // jim is a variable of type person, where person is a struct
jimPointer.updateName("Jimmy")
```



Call the receiver function (option 2)

```go
jim.updateName("Jimmy") // jim is a variable of type person, where person is a struct
```



The receiver function

```go
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
```



**jimPointer := &jim**

- is declaration of a pointer to a person
- pass the address of jim to jimPointer

***person** 

- syntax: is a star in front of the type
- is a type description - it means we are working with a pointer to a person

***p**

- syntax: is a star in front of a pointer
- is an operator - it means we want to manipulate the value the pointer is referencing



## Example with Slice - Exception Case

With slice, the behavior is different than with struct of the previous examples.

With slice, the slice is updated without having to use pointers.

```go
func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice) // the mySlice is modified without using pointers

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
```



# GO Routine

**Problem**: In the example below we can see that when the function checkLink is called, then the code will block until the code of the function is executed.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link) // blocking call
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

```





**WaitGroup**

Waits for the launched goroutine to finish

Package sync provides basic synchronization functionality

Functions used with a WaitGroup

- Add: Sets the number of goroutines to wait for (increases the counter by the provided number)
- Wait: Blocks until the WaitGroup counter is 0
- Done: Decrements the WaitGroup counter by 1. So this is called by the goroutine to indicate that it's finished



# Function Literal

https://dev.to/spindriftboi/function-literals-and-closure-in-go-2hgn
