# Go for programmers experienced with Python and C++
## Variables
```
var a \<type\> = \<value\> (zero-valued if value is not included)  
var a = \<value\> (Type Inference)  
a := \<value\> (Equivalent to var a = \<value\>)  
For constants, use const instead of var.  
"\_" is a blank identifier and can be used to ignore return values for functions
```

## For loop
Exactly like C/C++ for loops but without parentheses  
```
for \<init\>; \<break condition\>; \<increment\> { }  
```
There are no while loops in Go but for {} or for true {} will work like while true

### Range
range can be used to iterate over various data structures
Usage: for u, v := range data-structure { //body }  
**Array/slice**: u is the index and v is the element.  
**String**: u is the index and v is the rune int (int32 value for the ASCII character).  
**Map**: u is the key and v is the value of the key-value pair in map.  
**Channel**: u is the element and v is none.  

## If/Else
Exactly like C/C++ but without parentheses

## Switch-Case
Exactly like C/C++ but without parentheses in the switch expression and break statements are not needed

## Arrays and Maps
Both arrays and maps in Go are initialized using the builtin make() function  
Arrays can be sliced like Python but unlike Python all elements must have the same type and there is no step argument that lets you skip or reverse the order of elements  
To append individual elements into an array, use append(array, element1, element2, ...)  
To append elements of one array into another add an ellipsis (...) after the array  
  
### Examples:
```  
e := make([]string, 3) (Make a string of length 3)  
  
e := make(map[key-type]val-type)  
  
a := []string{"John", "Paul"}  
b := []string{"George", "Ringo", "Pete"}  
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"  
```
  
Maps work like dictionaries in Python but m[key] has an optional second return argument that returns true if the key-value pair exists.  
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
Returns types must be explicit like C++ but instead of starting with the return type, use the keyword "func" and put the return type after the function name and parameters  
Unlike C++ though, parameter types are declared *after* the parameter names and if multiple parameters have the same type, you can omit the type declaration up until the last parameter.  
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
