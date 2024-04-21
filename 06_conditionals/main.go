package main

import "fmt"

func main() {
	fmt.Println("hello world")
	x := 15
	y := 20

	if x <= y {
		fmt.Printf("%d value is lesser than %d\n", x, y)
	} else {
		fmt.Printf("%d value is greater than %d\n", x, y)
	}

	theme := "custom"

	switch theme {
	case "dark":
		fmt.Println("use dark theme")
	case "light":
		fmt.Println("use light theme")
	case "custom":
		fmt.Println("use custom theme")
	default:
		fmt.Println("use default theme")
	}
}
