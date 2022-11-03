package main

import "fmt"

func main() {
	//variable initialization with var and type
	var name string = "Oksir"
	fmt.Println(name)
	var isOk bool = true
	fmt.Println(isOk)
	var meta float32 = 23.53464535356456 //float32 store till decimal place 6
	var meti float64 = 23.53464535356456 //float64 stores more precise value (11 decimals)
	fmt.Println(meta)
	fmt.Println(meti)
	fmt.Printf("Type of meta variable is: %T \n", meta)
	fmt.Printf("Type of meti variable is: %T \n", meti)
	var x int
	fmt.Println(x) //initialized with 0;
	var str string
	fmt.Println(str) //initialized with ""
	var b bool
	fmt.Println(b) //initialized with false
	fmt.Println("Variables file initialized")

	//NOTE: cross type assignment is restricted, bool cannot later become string or int...

	// only var initialization. Initialization is necessary
	var username = "OkSIR"
	fmt.Println(username)

	//walarus operator of assignment.
	//Walarus operator are only allowed inside a method. Not available for use at global level.
	user := "oksir"
	fmt.Println(user)

	//NOTE
	//Naming of vaiable and functions by convention uses first smallcase char in private member and first upper case char as public member

	const pick int = 12322
	fmt.Println(pick)
}
