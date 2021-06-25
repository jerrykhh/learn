package main

//import "fmt" // Go lang IO Library, if not import it will error
import (
	"fmt"
	"math"
)

func main() {

	// Variable Delcare and data types

		var varString string = "initial"
		var varInt int = 10 // int int8 int16 int32 int64 #unsigned(u+): uint, ....
		var varBoolean bool = true
		var varFloat float32 = 10.0      // float32 float64
		var cal = varInt + int(varFloat) // it must change the datatype int()
		const finalVarData = "final"     // final variable
		varGenerics := 1                 // *Delcare Generics type [common use]

		fmt.Println(varString, varInt, varBoolean, varFloat, cal, finalVarData, varGenerics)
		// All variable must be used, if not it can't compiler

		// Array
			//delcare option1
			var langs1 [2]string
			langs1[0] = "Go"
			langs1[1] = "Python"

			//delcare option2
			langs2 := [2]string{
				"Go",
				"Python",
			}

			//delcare option3
			langs3 := []string{"Go", "Python"}

			fmt.Println(langs1, langs2, langs3)
			// out: [Go Python] [Go Python] [Go Python]

				//Array slices
				s := make([]string, 3) 		// create empty slice []
				s[0] = "a"
				s[1] = "b"
				s[2] = "c"
				fmt.Println("s:", s, "Length:", len(s))
				s = append(s, "d") 			// append data to current slice
				s = append(s, "e")
				c := make([]string, len(s))
				copy(c, s)          		// copy slice copy(dst, src)
				fmt.Println(c[2:5]) 		// out(index): 2, 3, 4
				fmt.Println(c[:5])  		// out(index): 0, 1, 2, 3, 4
				fmt.Println(c[2:])  		// out(index): 2, 3, 4, ...


		// Map/Dict(key:value)
		M := make(map[string]int)
		M["jerry"] = 1
		fmt.Println("Length:", len(M))
		delete(M, "jerry")			// delete data in Map delete(MAP, key)


		// Pointer
		var pointerString *string = &varString // & get the variable memory address
		fmt.Println(pointerString)             // Memory Address
		fmt.Println(*pointerString)            // Data of Memory Address

	// Constant
	//math
	fmt.Println(math.Pi)
	fmt.Println(math.Sin(math.Pi / 6)) // .Cos(x float64) .Tan(x float64)
	fmt.Println(math.Pow(7, 2))        // out: 49
}
