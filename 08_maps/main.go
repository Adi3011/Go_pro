package main

import "fmt"

func main() {
	// email := make(map[string]string) // map with key and value as strings map[string]string means map[key type]value type

	// // assign values
	// email["aditya"] = "aditya@gmail.com"
	// email["anurag"] = "anurag@gmail.com"
	// email["amit"] = "amit@gmail.com"

	// or we can do something like below without using make
	email := map[string]string{"aditya": "aditya@gmail.com", "anurag": "anurag@gmail.com", "amit": "amit@gmail.com"}

	email["charlie"] = "charlie@gmail.com"

	fmt.Println(len(email))
	fmt.Println(email["aditya"])

	// delete from map
	delete(email, "aditya")
	fmt.Println(email["aditya"])
	fmt.Println(email)
}
