## Hands-on exercise #1
1. Start with this code (https://play.golang.org/p/3W69TH4nON).
2. Instead of using the blank identifier, make sure the code is checking and handling the error.

**solution:** https://play.golang.org/p/tn8oiuL1Yn 

## Hands-on exercise #2
1. Start with this code (https://play.golang.org/p/9a1IAWy5E6).
2. Create a custom error message using `fmt.Errorf`.

**solutions:**
- https://play.golang.org/p/HugU4HJEEO 
- https://play.golang.org/p/NII-lmGasj 
- https://play.golang.org/p/Vo5kIoR-sG 

## Hands-on exercise #3
1. Create a struct `customErr` which implements the `builtin.error` interface. 
2. Create a func `foo` that has a value of type error as a parameter. 
3. Create a value of type `customErr` and pass it into `foo` (hint: https://play.golang.org/p/L5QWV8-p11).

**solution:** https://play.golang.org/p/ixeowY2fd2 

(Other possibilities: **assertion:** https://play.golang.org/p/pbl2kCYsM0 or **conversion:** https://play.golang.org/p/1ldiBdkdzA)

## Hands-on exercise #4
1. Start with this code (https://play.golang.org/p/wlEM1tgfQD) and use the `sqrt.Error` struct as a value of type error. 
2. Use these numbers:
   - lat "50.2289 N"
   - long "99.4656 W"

**solution:** https://play.golang.org/p/nsRxbDfkCh 