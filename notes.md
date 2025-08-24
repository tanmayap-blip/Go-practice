 # Go Notes 

Go (created at Google in 2009) is a compiled, statically typed, garbage-collected language designed to be:

Simple (small syntax, no hidden magic).

Fast (compiled to machine code, near C/C++ speed).

Concurrent (built-in goroutines & channels).

Practical (standard library covers most real-world backend needs).

# Exported Names 

In Go, a name (variable, function, type, constant, struct field, etc.) is said to be exported if it starts with an uppercase letter.

Exported name → Visible outside the package (public).

Unexported name → Visible only inside the package (private).
There are no public, private, or protected keywords like in Java/C++ — just capitalization!

In go we need to keep the track of exported and unexported variable and function names or nay types of names , as in go we everything in  package , so in case of unexported if they are same then will create error for the inside of the package and in case of exported names they will create problmes outside the or inside the package 

# functions 

functions are the heart of Go (since Go doesn’t have classes like Java or C++, you use functions + structs + interfaces to build programs).

Multiple Return values - in functions the first things are the set of set of parameters then if we want to return mutiple values so we can add the type list , which will tell the return of multiple values type 

Named return values - in this if we want to name the return values and then we will update the values and directly  use the return then all the return values will be returned 

Variadic functions - in this kind of functions we can pass any number of arguments into the funcion, the number of arguments passed is like nums ... is passed ike a slice 

# closures 

A closure is a function value that “remembers” the variables from the scope where it was created, even if that scope is gone.


# pointers

Pointers are used to store the memory address of the variables 
With pointers we can implement the pass by referce functionality , such that the refernce of the variable or the data structure is passed and we can directly do operations inside any of the function 
Go doest not have pointer arithmetic 
it does not have manual memory allocation 
we can use the new () for allocating memory to any variable 

# structs 

Structs are used to create  a custom data type by cobining multiple fields of same and different types 

we can copy the instance of the struct in two ways one is the direct way , in this the coped instance is different means if we change anything then it will not chnage in the original instance , and there is a referce way in which we copy the struct instance using reference and by this way if we change anything using the copied instance then we the original instance is also changed 

# methods

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

Go already has functions, but methods allow us to attach behavior directly to structs. This gives encapsulation, makes code feel more object-oriented, enables structs to satisfy interfaces, and provides cleaner syntax like p.Greet() instead of Greet(p).

# Interface 

An interface in Go is a type that specifies a set of method signatures (a contract).
If a type (struct or custom type) has all the methods required by the interface, then it automatically implements that interface — no implements or extends keyword is needed.

They enable polymorphism in Go.

There are diffferent types of buit in interface like Stringer and error 

