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

# concorruency in Go 

It looks like the different processes are running concorruently , but actually they are runnig by a single processing unit and by context switching and the operation becomes so fast that it looks like they are running concurrently 

# Go Routine

It is one of the method for executing the concurreny in go lang 

Every process which run concurrently in golang i called goroutine 
it is  light weighted thread 
creation cost of go routine is smaller than thread 
every program has atleast one go routine named as main function 
when main is terminated then all the go rutines will be terminated 

we will be using time .Sleep for giving time to the goroutine to start executing 

go routines use the staack for the excution , but we can increase or decrease the memory used by the go routine , but in case of thread we cannot do it 

# Channels 

channel is a communication pipeline that allows multiple goroutines to safely send and receive data from each other

when there is  dependency of  one go routine in other then we need to use channels so that they can communicate 


Imagine two people, Goroutine A and Goroutine B, who need to pass a message to each other. They don't have a shared desk, so they use a special mailbox.

Sending: Goroutine A writes a message and puts it in the mailbox.

Receiving: Goroutine B opens the mailbox and takes the message out.

The channel is this mailbox. It ensures that the message is received by only one goroutine and that the data transfer is safe.

in case of unbuffered channels that we create normally ,  if the receiver is not set , means the chanel is not used by any of the go routine than the main function blocks the code at that point of time and the compiler then gets to know that there is only one go routine and it is also blocked  therefor it enters inside the deadlock condition 

in case of unbuffered channels ,  both the sender and receiver should be present when we use the channels , else there will be a deadlock condition 

ok status is used to check the check if the channel is closed or not 

buffered channel - to make the channel buffered we need to set the size of the channel to some fixed value 


There are two types of channel , When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.

1) Bidirectional channel - when the channel is not specified anything means it can both send and receive data 
 ex- func sending(ch chan string){}
2) Unidirectional Channel - when the channel is mentioned with either send only or receive only as its type , send only is intialized as ch chan<- string, and the receive only is intialized as ch <-chan string
   ex- func sending(ch chan<- string){}  means you can only update the channel with the data 
3) ex- func receiving(ch <-chan string){} means you can only read from the channel not possible to send anything 


# Channel Synchronization 

When one goroutine needs to pass something to another, they both must be "present" for the exchange to happen. If the sender gets there first, it waits for the receiver. If the receiver gets there first, it waits for the sender. This waiting and signaling is what channel synchronization is all about. It guarantees that certain operations happen in a predictable order, preventing chaos and ensuring your program works correctly.

Channel synchronization relies on the blocking behavior of channels.

Blocking on Send: A goroutine sending a value to an unbuffered channel will block (pause) until another goroutine is ready to receive that value. This is a synchronization point.

Blocking on Receive: A goroutine receiving a value from an empty channel will block until a value is sent to it. This is another synchronization point.


in case of only one go routine need to be waited we can do that using the time.Sleep 
but when we need to wait for multiple go routines to finish we can use wait Group 

there are three function in wait group add for assigfning number of go routines , wait for making the code wait until all the go routines are not completed , done for marking one go routine completed , we basically create a variable for wait group with sync.waitgroup, then we mark the number of gr then we pass the refernce of wait group for the done in the go routine using defer , then at the end of the code we will do the wait 
 

# Select 

Select statement is basically used for waiting on multiple channel operations at once .It chooses which channel operation to run based on which one is ready first.

A select statement contains a set of case clauses and an optional default clause.

case Clause: Each case corresponds to a channel operation (either sending or receiving). The select statement evaluates all of these cases.

Ready Channels: The select statement will choose a case whose channel is ready to proceed.

For a receive operation (<-ch), the channel is ready if it has a value in its buffer or a sender is available.

For a send operation (ch <- value), the channel is ready if it has space in its buffer or a receiver is available.

No Ready Channels: If none of the channels are ready, the select statement blocks until one of them becomes ready.

The default Clause: The default clause is optional. If it's present, and no other case is ready, the select statement will execute the default block immediately without blocking.


-> Non blocking channels

In Go, all channel operations are blocking by default. However, you can achieve non-blocking channel operations by using a select statement with a default clause.

How It Works
The select statement with a default clause gives you a way to check a channel's status without getting stuck.

Blocking Operation: If you try to send to or receive from a channel, and it isn't immediately ready (e.g., the channel is empty or full), your goroutine will pause and wait. This is a blocking operation.

Non-Blocking Operation: By adding a default clause to a select statement, you tell the program: "If the channel isn't ready right now, don't wait. Just run this other code instead." This allows your program to keep moving without being halted.

This is extremely useful when you want to avoid deadlocks or to periodically check for new data without dedicating a goroutine to wait indefinitely.


-> timeouts - there is a feature of timeouts , in this we can assign a timeout in the select statements , such  that if none of the case is ready till the timeout the timeout case will be used 

``` select {
case res := <-c1:
fmt.Println(res)
case <-time.After(1 * time.Second):
fmt.Println("timeout 1")
} 
```


# Worker Pool

A worker pool is a design pattern in concurrent programming that involves a fixed number of goroutines (workers) that are created once and then continuously wait for and process tasks from a shared queue (a channel).

A worker pool is a design pattern in concurrent programming used to manage a fixed number of goroutines (workers) to execute a large number of tasks. Instead of creating a new goroutine for every single task, you create a limited pool of workers that sit and wait for work to be assigned to them

A typical worker pool in Go consists of three main parts:

A Jobs Channel: This is a channel that acts as a queue. Producers (the main goroutine or other parts of the program) send tasks to this channel.

Worker Goroutines: These are the goroutines that make up the pool. They are launched at the start and continuously read from the jobs channel, process the tasks, and then wait for the next one.

A Results Channel: (Optional but common) This is another channel where workers can send their results back to a collector.

# Atomic Counter 

An atomic counter is a special type of counter that is guaranteed to be updated safely in a concurrent environment. It's used to increment or decrement a number without any chance of a race condition.In simple terms, an "atomic" operation is an action that happens as a single, indivisible unit. It either completes entirely or doesn't happen at all. No other goroutine can interrupt it.



The Problem: Race Conditions
Without atomic operations, a simple counter in a concurrent program can lead to unpredictable results. This is because a standard increment operation (counter++) is not atomic; it's actually three separate steps:

Read the current value of the counter.

Add 1 to that value.

Write the new value back to the counter.



# mutex 

A mutex (short for mutual exclusion) is a synchronization primitive used to protect shared data from being accessed by multiple goroutines at the same time. It acts like a lock that ensures only one goroutine can access a critical section of code at any given moment.

How Mutexes Work
A mutex provides two primary methods:

Lock(): This method acquires the lock. If the lock is already held by another goroutine, the current goroutine will block (pause) until the lock is available.

Unlock(): This method releases the lock, allowing other waiting goroutines to acquire it.

By placing Lock() and Unlock() calls around a shared resource, you create a "critical section" of code that is guaranteed to be executed by only one goroutine at a time.



# Panics 

A panic in Go is a special type of error that stops the normal execution of a program. It's an unrecoverable runtime error, typically used for situations where a program cannot continue safely, such as an out-of-bounds array access, a nil pointer dereference, or a division-by-zero error. When a function panics, it immediately stops executing, and the Go runtime begins to unwind the call stack, running any deferred functions along the way before terminating the program and printing a stack trace.

recover is a function which is used to catch panics , there fore it is needed to be passed as the defer , in the recover it will return the value passed in panic 


# Unit Testing 
There are different kind of testing : - unit , integration, acceptance and etc ,
Unit means a small peice of code often refferd to a functions 
unit testing means testing each piece of code in isolation , independent of other codes 

benefits of unit testing :-
1) isolation 
2) improved code quality 
3) code coverage 
4) regression detection - by mistake changes detection 
5) refactoring and maintainance

 -> Mocking

 Mocking in unit testing is the process of creating a fake, stand-in version of an external service or a complex dependency that your code relies on.

Instead of your code talking to a real database, a remote API, or a file system during a test, it talks to a mock object that you create. The mock object behaves exactly as you tell it to—it can pretend to return specific data, simulate an error, or just record that a certain function was called.


# Modules 

Modules is a collection of packages stored in a file tree with a go.mod file at its root

# Http 

-> Status Code 
Info - 100-199
Success- 200- 299
Redirection - 300-399
Client Error- 400-499
Server Error- 500- 599


ServeMux is Gos built in request router , which handles the request calls implicitly 





Read about single flight, why we use , advantages and disadvantages 

visualize panics , and make sure it should not be shown on server , basically handle panics 


gin :- 

normal server
midllware
conversation between middleware
api creation 


gorm :-
sql download
migration 
use gorm foe all of this 
goose

testing 

suite



