# Go for programmers experienced with Python and C++
## Variables

`var a <type> = <value>` (zero-valued if value is not included)  
`var a = <value>` (Type Inference)  
`a := <value>` (Equivalent to var a = \<value\>)

For constants, use `const` instead of `var`.  
`_` is the blank identifier and can be used to ignore return values for functions.


## For loop
Exactly like C/C++ for loops but without parentheses  

`for <init>; <break condition>; <increment> { //body }`

There are no while loops in Go but for {} or for true {} will work like while true.

### Range
Range can be used to iterate over various data structures.
Usage: `for u, v := range data-structure { //body }`
**Array/slice**: u is the index and v is the element.  
**String**: u is the index and v is the rune int (int32 value for the ASCII character).  
**Map**: u is the key and v is the value of the key-value pair in map.  
**Channel**: u is the element and v is none.  

## If/Else
Exactly like C/C++ but without parentheses.

## Switch-Case
Exactly like C/C++ but without parentheses in the switch expression and break statements are not needed.

## Arrays and Maps
Both arrays and maps in Go are initialized using the built-in `make()` function.
Arrays can be sliced like Python, but unlike Python, all elements must have the same type. There is no step argument that lets you skip or reverse the order of elements.  
To append individual elements into an array, use `append(array, element1, element2, ...)`.
To append elements of one array into another, add an ellipsis (...) after the array.
  
### Examples:
```  
e := make([]string, 3) (Make a string of length 3)  
  
e := make(map[key-type]val-type)  
  
a := []string{"John", "Paul"}  
b := []string{"George", "Ringo", "Pete"}  
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"  
```
  
Maps work like dictionaries in Python but `m[key]` has an optional second return argument that returns true if the key-value pair exists.  
Example: `_, retval := map[nonexistent-key] (retval is false)`

## Defer
Will execute a function after the function the defer call is inside returns.

**Example**:
```
func main() {
  defer fmt.Println("world")  // prints second (after main() finishes)
  fmt.Println("hello")        // prints first
}
```

## Functions
Return types must be explicit like C++, but instead of starting with the return type, use the keyword `func` and put the return type after the function name and parameters.
Unlike C++ though, parameter types are declared *after* the parameter names, and if multiple parameters have the same type, you can omit the type declaration up until the last parameter.  

**Examples**: 
``` 
func fib(n uint) uint {  
  if n < 1 return n  
  else return fib(n - 1) + fib(n - 2)  
}  
func sum(a, b, c int) int {  
  return a + b + c  
}
```

Functions in Go can also support multiple return values. The type for each return value can be specified in a set of parentheses after specifying the function's arguments. If you do not want certain values, you can obtain a subset of return values by using the blank identifier.

**Example**:
```
func ispos(a int) (string, int) {
    if a >= 0 {
      return "Yes", 1
    } else {
      return "No", 0
    }
}

func main() {
    res, _ := ispos(0)
    fmt.Println(res)
}
```

It is possible for Go functions to take in a variable number of inputs, too - use a container variable, an ellipsis, and a type within the argument parentheses. We call these variadic functions.
When calling them, one can use individual arguments or multiple arguments provided by a slice. To do this with a slice, you should follow the format `func(slice...)`.

**Example**:
```
func prod(vals ...int) {
    fmt.Print(vals, " ")
    tot := 0
    for _, val := range vals {
        tot *= val
    }
    fmt.Println(tot)
}

func main() {
    prod(0, 1, 1)
    vals := []int{2, 3, 5, 8}
    prod(vals...)
}
```

Also possible in Go is anonymous functions, which allows us to declare functions inline without needing to name them. When we return from a function defined anonymously in another function's body, the returned function forms closure with the variables used inside it.

**Example**:
```
func nexteven() func() int {
    a := 0
    return func() int {
        a += 2
        return a
    }
}

func main() {
    next := nexteven()  // next has its own a value, which is updated on every next() call
    fmt.Print(next())
    fmt.Print(next())
    fmt.Println(next())
}
```

Recursion works just like C/C++ but with Go's conventions.

## Pointers

Same idea as C++ pointers - star (\*) dereferences, ampersand (&) gives memory address.

## Structs

Similar to C/C++, though different syntax is present. Structs can be changed.

**Example**:
```
# Struct
type employee struct {
    name string
    department int
}

# Constructor
func newEmployee(name string) *employee {
    e := employee{name: name}
    e.department = 7
    return &p
}

# How to create a new struct
    fmt.Println(employee{"Gordon", 4}
    fmt.Println(newEmployee("Joe"))

# You can name fields while initializing a struct; if a field is not given a value,
# it will be zero-valued
    e := employee{name: "Bob", department: 7}

# You can access struct fields with a period
    fmt.Println(e.name)
```

## Goroutine

Executes functions concurrently with calling functions

`go compare(2, 2)`
