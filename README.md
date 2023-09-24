# ARCANE Programming Language

Arcane is a programming language designed to support a range of essential features including functions, arrays, closures, objects, and built-in functions like len, log, push, first, last, and rest.

## Table of Contents

- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Running Programs](#running-programs)
- [Language Features](#language-features)
  - [Functions](#functions)
  - [Arrays](#arrays)
  - [Closures](#closures)
  - [Objects](#objects)
  - [Built-in Functions](#built-in-functions)

## Getting Started

### Installation

To use Arcane, follow these steps:

1. Clone the repository or download the source code.
2. Compile the source code using `go run main.go` or `go build` and run the `arcane.exe`.

### Running Programs

Once you have that set up, you can start coding:

## Language Features

### Functions

Arcane supports functions, allowing you to define and call functions within your code. Here's a basic example of a function:

`let sum = fn(a, b) {
return a + b;
}`

`let result = sum(10, 5);`
`log(result); # Outputs: 15`

### Arrays

Arrays can be used to store and manipulate collections of data. Here's an example of creating and using an array:

`let users = ["John", "Jane", "Jack", "James"];`

`log(users); #Outputs: ["John", "Jane", "Jack", "James"]`

`log(len(users)); #Outputs: 4`

`log(first(users)); #Outputs: "John"`

`log(last(users)); #Outputs: "James"`

`log(push(users, "Julian")); #Outputs: ["John", "Jane", "Jack", "James", "Julian"]`

### Objects

Objects allow you to define custom data types with properties. Here's an example of defining and using an object:


`let people = [{"name": "Jack", "age": 24}, {"name": "Sparrow", "age": 99}];`

`let getName = fn(person) {person["name"];};`

`getName([0]); #Outputs: "Jack"`
`getName([1]); #Outputs: "Sparrow"`

### Closures

Closures allow you to create functions with access to variables from their surrounding scope. Here's a simple closure example:

`let newAdder = fn(a, b) {
fn(c) { a + b + c };
};`

`let adder = newAdder(1, 2); # This constructs a new `adder` function:`

`adder(8); #Outputs 11`

### Built-in Functions

Arcane provides several built-in functions for common operations:

`myArray = [1, 2, 3];`
`log(len(myArray)); # Outputs: 3`

`log("Logging a message."); # Outputs: Logging a message.`

`push(myArray, 4);`

`log(myArray); # Outputs: [1, 2, 3, 4]`

`log(first(myArray)); # Outputs: 1`
`log(last(myArray)); # Outputs: 4`
`log(rest(myArray)); # Outputs: [2, 3, 4]`
