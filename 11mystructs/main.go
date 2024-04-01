package main

import "fmt"

func main() {
	fmt.Println("Welcome in Structs")

	ramgopal :=  User{"Ram Gopal", "ramgoal@gmail.com", true, 20}
	fmt.Println(ramgopal)
	fmt.Printf("Ram Gopal full details are : %+v\n", ramgopal)
	fmt.Printf("Name is %v and email is %v.\n", ramgopal.Name, ramgopal.Email)


}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}
