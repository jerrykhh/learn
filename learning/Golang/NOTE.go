package main

//import "fmt" // Go lang IO Library, if not import it will error
import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

func init() {							// Golang is able to init(), it will run it before main()
	fmt.Println("init...")
}

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

  	// if-else

		var mark float32 = 100
		if mark > 50 {								// it not need write ()
			fmt.Println(">50 mark")
		} else if mark > 39 {
			fmt.Println(">=40 && <=50 mark")
		} else {
			fmt.Print("Fail")
		}

  	// switch

		mark = 50
		switch {									// Also, it can switch(variable){}
			case mark > 50:
				fmt.Println(">50 mark")
			case mark > 39:
				fmt.Println(">=40 && <=50 mark")
			default:
				fmt.Print("Fail")
		}

	// Loop
		var count int = 1
		for count <= 5 {							// like while loop
			fmt.Println("count:", count)
			count++
		}

		for i := 1; i <= 5; i++{					// basic for loop; it cannot wrtie: var i int = 1
			fmt.Println("index ", i)
		}

		count = 0
		for {										// Infinite loop
			fmt.Println("count:", count)
			if count == 5{
				break								// break or continue can use
			}
			count++
		}

		// range (Array)
			std_names := []string{"Jerry", "Marry", "Jason"}
			for index, name := range std_names{		// it cannot: for name:= range std_names{}
				fmt.Println("index:", index, ", name:", name)
			}

		//range (Map)
			students := map[int]string{1: "Jerry", 2: "Marry", 3: "Jason"}
			for key, value := range students {
				fmt.Println("studentId:", key, ", name:", value)
			}
			for value := range students {			// Map loop all value only
				fmt.Println("name:", value)
			}

  	// Function
		fmt.Println(intPlus(1, 1))
		fmt.Println(intPlus2(1, 1))
		result := 0
		plus(&result, 1, 1)
		fmt.Println(result)
		fmt.Println(sum(1, 1))
		fmt.Println(returnMultiValue(1, 3)) 		// output: 4 -2 3 0
			// Anonymous Function
			ana_func := anonymous_func(1,1)
			fmt.Println(ana_func())
			ana_func1 := func (x, y int) int {
				return x + y
			}
			fmt.Println(ana_func1(1,1))				// output: 2

	// Struct
	type student struct {
		id int
		name string
	}

	var student_list []student					// create Object List
	count = 1
	for count <= 5 {
		student_list = append(student_list, student{id: count, name: "jerry"})
		count++
	}

	for _, std := range student_list {
		fmt.Println(std.id, std.name)
	}

	// Interface main

		GetObjectName(&Apple{name: "happy"})
		GetObjectName(&Person{name: "handsome"})


	// Error handle / Exception main

		fmt.Println(divide(1, 0))
		result, err := cus_divide(1, 0)
		if err != nil{
			fmt.Println(err)
		}else {
			fmt.Println(result)
		}


	// Goroutines & Channel main

		user1_chat_channel := make(chan message)
		user2_chat_channel := make(chan message)		// channel can use other data type / struct
		user1 := User{name: "Jerry"}
		user2 := User{name: "Terry"}
		defer close(user1_chat_channel)					// 'defer' will run all program then will run, also it like push to defer stack
		defer close(user2_chat_channel)

		go message_channel(user1, user1_chat_channel)	// go mean async function
		go message_channel(user2, user2_chat_channel)

		user1_chat_channel <- message{user: user2, mes: "Hello world"}
		time.Sleep(time.Second)
		user2_chat_channel <- message{user: user1, mes: "Hello"}
		time.Sleep(time.Second)
		user1_chat_channel <- message{user: user2, mes: "Good morning"}
		time.Sleep(time.Second)
		user2_chat_channel <- message{user: user1, mes: "Good night"}
		time.Sleep(time.Second)

		/*
			output:
			Sender: Jerry , Receiver: Terry, Message: Hello world
			Sender: Terry , Receiver: Jerry, Message: Hello
			Sender: Jerry , Receiver: Terry, Message: Good morning
			Sender: Terry , Receiver: Jerry, Message: Good night
		*/

		fmt.Println("End")

}

// Function

	func intPlus(x int, y int) int {				// Delcare func funcName(Param) return type {}
		return x + y
	}

	func intPlus2(x, y int) int {					// if all params is same data type, it write the last only is ok
		return x + y
	}

	func plus(value* int, x int, y int) {			// if not return type (void), it not need to write the type
		*value = x + y
	}

	func sum(nums ...int) int {						// dynamic params
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return sum
	}

	func returnMultiValue(x, y int) (int, int, int, int) {		// it can return many value
		return x+y, x-y, x*y, x/y
	}

	func anonymous_func(x, y int) func() int {					// Anonymous Function
		return func () int {
			return x + y
		}
	}

// Interface (it can use to delcare the Getter/ Setter)
	type Object interface {						// abstract func
		getName() string
		setName(name string)
	}

	type Person struct {						// Person struct and Getter& Setter
		name string
	}

	func (p* Person) getName() string {
		return p.name + " person"
	}

	func (p* Person) setName(name string) {
		p.name = name
	}
	type Apple struct {							// Apple struct and Getter& Setter
		name string
	}

	func (a* Apple) getName() string {
		return a.name + " apple"
	}

	func (a* Apple) setName(name string) {
		a.name = name
	}

	func GetObjectName(obj Object) {				// Polymorphism
		fmt.Println(reflect.TypeOf(obj), obj)
		fmt.Println(obj.getName())
	}

// Error handle / Exception
	/*
		Golang implemented the error struct (import "errors")
		type error struct {
			Error() string
		}
	*/

	func divide(numerator int, denominator int) (int, error) {
		if denominator == 0 {
			return -1, errors.New("Denominator cannot be 0")
		}
		return numerator/denominator, nil
	}

	// Customize Error / Exception

		type DivisionByZeroError struct {						// Customize Error struct
			num int
			desc string
		}

		func (err* DivisionByZeroError) Error() string {
			return fmt.Sprintf("%d -> %s", err.num, err.desc)
		}

		func cus_divide(numerator int, denominator int) (int, error){
			if denominator == 0 {
				return -1, &DivisionByZeroError{denominator, "DivisionByZeroError"}
			}
			return numerator/denominator, nil
		}

	// Goroutines & Channel

		type User struct{
			name string
		}

		type message struct {
			user User
			mes string
		}

		func message_channel(p_user User, p_chan chan message){
			fmt.Println(p_user.name, "online")
			for{
				value, ok := <- p_chan									// it will return data, (true or false)
				if ok {
					fmt.Println("Sender:", p_user.name, ", Receiver:", value.user.name + ",", "Message:", value.mes)
				} else {
					break
				}
				time.Sleep(time.Second)
			}
		}
