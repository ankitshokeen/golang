package main

import "fmt"

func function() {
	fmt.Println("welcome to functions...")

	// passing arguments as integer and strings
	fmt.Println(str("hello ", "ankit"))
	fmt.Println(intg(20, 12, 20))

	// passing slice
	fmt.Println(slicess("hello ", "everyone ", "i`m ", "go ", "developer "))

	fmt.Println(uData(0, "", "", 0))

}

func str(s1 string, s2 string) string {
	return s1 + s2
}

func intg(n1 int, n2 int, n3 int) int {
	return n1 - n2 + n3
}

func slicess(greet ...string) string {
	grt := ""

	for _, v := range greet {
		grt += v
	}

	return grt
}

// passing multiple return arguments
func uData(id int, name string, email string, phone int) (int, string) {
	if id != 0 && name != "" && email != "" && phone != 0 {
		return 200, "OK"
	} else {
		return 500, "ERROR"
	}
}
