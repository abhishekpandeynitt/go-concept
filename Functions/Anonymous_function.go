//Anonymous function is a function that we use for very short calulaion within main
//it does not have any function name, we directly call it with func keyword with givn arguments

package main

import "fmt"

func main() {
	x := func(a, b int) int {
		return a + b
	}

	fmt.Printf("DataType of x is %T\n", x)
	fmt.Printf("Value of x is %d", x(20, 30)) //we are calling by giving value aruguments here.(int)

	y := func(a, b int) int {
		return a * b
	}(10, 20)

	fmt.Printf("DataType of x is %T\n", y)
	fmt.Println("Value of x is", y) //calling y with directly aissned value

}
