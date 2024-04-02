package main

import "fmt"

func main() {
	// Welcome message
	fmt.Println("Welcome in methods")

	ramgopal :=  User{"Ram Gopal", "ramgoal@gmail.com", true, 20}
	fmt.Println(ramgopal)
	fmt.Printf("Ram Gopal full details are : %+v\n", ramgopal)
	fmt.Printf("Name is %v and email is %v.\n", ramgopal.Name, ramgopal.Email)

	// Access method
	ramgopal.GetStatus()

	// Print altered email
	ramgopal.NewMail()

}


type User struct {
	Name string
	Email string
	Status bool
	Age int
}


// Pass struct in function and create method
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}


// Alter email address
func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Updated Email is: ", u.Email)
}
