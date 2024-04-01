package main

import "fmt"

func main(){
	fmt.Println("Welcome in Maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("Full form of JS is: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of languages after delete RB : ", languages)


	// Print map by use of loop 
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}