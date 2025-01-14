Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
>concurrency refers to a system's ability to handle multiple tasks at once, but not necessarily simultaniously. Parallelism is a subset of concurrency, where tasks are actually handled simultaniously (typically on separate processors).  

What is the difference between a *race condition* and a *data race*? 
> a race condition refers to behaviour where the system output depends on the timing or order of events. A data race is a race condition where to threads are accessing the same memory and at least one of these are attempting a write to that location. 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler is a mechanism within an OS that decides what process or thread to run next. The scheduler keeps a queue of ready to tun processes and determines which one to run based on its own policy (algorithm). 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> We should use multiple threads to perform concurrent tasks. This allows multiple tasks to run simultaneously and solve problems that require multitasking. We can use this if we want a process to handle one task whilst others handle background tasks. We can also use multithreading for resource sharing.  

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> Fibers are a lightweight abstraction for concurrent execution. They are managed by the program itself, rather than the OS, and are usually cooperative. Coroutines are similar in that they allow functions to pause execution and resume later - opening up for asynchronous programming. We would rather use fibers and/or coroutines in instances where we dont want to deal with the full memory and/or context switch overhead of threads. This could, for example, be when dealing with multiple small tasks or IO. 

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Paralell execution becomes easier, as well as problems where resource sharing is critical. On the other hand, the programmer now has to deal with concurrency problems like race conditions, synchronous execution etc. 

What do you think is best - *shared variables* or *message passing*?
> They have different uses. Shared variables require less overhead, and are thus nice when shared states and/or efficiency is key. Message passing is better when we have complex concurrency and we dont want to deal with the pitfalls of shared variables - such as synchronisation and consistency between threads. 


