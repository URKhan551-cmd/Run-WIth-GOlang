package main  // declares executables
import fmt

func main(){
  fmt.Println("Hello, World!")
}

// in go source file belongs to a package
// package is a way to organize Source file and their functoinality
// fmt is a standard library packages provides formated input and output fuunctionality.
// in fmt the Println   here the p is capital because it asume it comes thrugh import 
// in go executables the runtime automatically calls main()
// we dont need to explicitly define main()  to start program

// go run main.go

func main(){
   name := "umair"  // short variables decleration  autmatically infers type of that variable. type= string
  age := 26  // type infers = int

  fmt.Println("My name is:", name)
  fmt.Println("i am", age, "old") 
}

// short variable decleration :=    Go declare variable annd asssign a value and infer its type.
// assignment operator   =    

name := "umair"
name = "ubaid"    // assigned new val to the first declared variable.

func main(){

  age := 36    // declared a var and initialized it
  age = 67

  fmt.Println(age) // 67
}


:= declaration + initialization
= assignment to an existing situation.


// DECLARING MULTIPLES VARIABLES.
name, age := "umair", 56



// ASSIGNING MULTIPLE VARIABLE.
name := "umaier"
age := 35

name, age = "uzhsj", 67




// RULE 1:   we cannot use :=  outside of a func
// short variables are only permitted inside a funcc
name := "umaie"  // invalid
var name  = "umaie" // valid
func main(){

  fmt.Println(name)
}


RULE :2;  we cannot redeclare an existing variable in the same scope 
func main(){
  age := 56
  age := 67 // invalid 
}

// Rule 3:  we can use  :=  with an existing var if at least one var is new
func main(){
  age := 45

  age, name := 67, "umair"  // valid   one name get declared and age assigned new one
}


// VAR
var age int = 67

// we can also use without value consider as 0
var age int


