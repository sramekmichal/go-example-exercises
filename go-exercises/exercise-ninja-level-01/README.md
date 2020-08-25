## Hands-on exercise #1
1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”.
   - 42
   - "James Bond"
   - true
2. Now print the values stored in those variables using
			a) a single print statement
			b) multiple print statements

**solution:** https://play.golang.org/p/yYXnWXIQNa

## Hands-on exercise #2
1. Use `var` to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE):
   - identifier “x” type `int`
   - identifier “y” type `string`
   - identifier “z” type `bool`
2. in `func main`
   - print out the values for each identifier
   - The compiler assigned values to the variables, what are these values called?

**solution:** https://play.golang.org/p/jzHwSlles9 

## Hands-on exercise #3
Using the code from the Hands-on exercise #2:
1. At the package level scope, assign the following values to the three variables:
   - for x assign 42
   - for y assign “James Bond”
   - for z assign true
2. in `func main`:
   - use `fmt.Sprintf` to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
   - print out the value stored by variable “s”

**solution:** https://play.golang.org/p/QFctSQB_h3  

## Hands-on exercise #4
For this exercise:
1. Create your own type - have the underlying type be an `int`.
2. Create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
3. in `func main`
   - print out the value of the variable “x”
   - print out the type of the variable “x”
   - assign 42 to the VARIABLE “x” using the `=` OPERATOR	
   - print out the value of the variable “x”

**solution:** https://play.golang.org/p/snm4WuuYmG 

*PS: FYI - nice documentation and new terminology “underlying type:” https://golang.org/ref/spec#Types*

## Hands-on exercise #5
Building on the code from the previous Hands-on #4:
1. At the package level scope, using the `var` keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”

eg:	
```
type chleba int
var x chleba
var y int
```

2. In `func main`:
   - this should already be done:
		- print out the value of the variable “x”
		- print out the type of the variable “x”
		- assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
		- print out the value of the variable “x”
   - now do this:
		- now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
		- then use the `=` operator to ASSIGN that value to “y”		
		- print out the value stored in “y”		
		- print out the type of “y”

**solution:** https://play.golang.org/p/cj8RrYgBOD