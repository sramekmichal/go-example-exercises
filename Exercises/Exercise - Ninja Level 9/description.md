## Hands-on exercise #1
1. In addition to the main goroutine, launch two additional goroutines.
    - Each additional goroutine should print something out.
2. Use wait-groups to make sure each goroutine finishes before your program exists.

**solution:** https://github.com/GoesToEleven/go-programming 

## Hands-on exercise #2
This exercise will reinforce our understanding of method sets:
1. Create a type `person` struct.
2. Attach a function `speak` to type person using a pointer receiver:
   - *person
3. Create a type `human` interface:
   - to implicitly implement the interface, a human must have the speak method.
4. Create func `saySomething`
   - have it take in a human as a parameter
   - have it call the speak method
5. Show the following in your code.
   - you CAN pass a value of type *person into saySomething
   - you CANNOT pass a value of type person into saySomething

Here is a hint if you need some help: https://play.golang.org/p/FAwcQbNtMG 

   Receivers  | Values 
   :---------:|:------:
   (t T)      | T and *T
   (t *T)     |*T   

**solution:** https://github.com/GoesToEleven/go-programming

## Hands-on exercise #3
1. Using goroutines, create an incrementer program:
    1. to have a variable to hold the incrementer value
    2. to launch a bunch of GoRoutines
        - Each GoRoutine should:
            1. Read the incrementer value
                - Store it in a new variable
            2. Yield the processor with runtime.Gosched()
            3. Increment the new variable
            4. Write the value in the new variable back to the incrementer variable
2. Use wait-groups to wait for all of your goroutines to finish.
3. The above will create a race condition.
4. Prove that it is a race condition by using the -race flag

Here is a hint if you need some help: https://play.golang.org/p/FYGoflKQej 

**solution:**: https://github.com/GoesToEleven/go-programming 
