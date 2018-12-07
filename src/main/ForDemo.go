package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i,"\t")
	}

	message := "This is error, and the error message is you couldn't find the type fo data";

	//message = "中华人民共和国"

	for i := 0; i < len(message); i++ {
		fmt.Printf("%c",message[i])
	}

	fmt.Println()

	var a int = 1

	LOOP:
		if a < 20 {
			fmt.Print(a,"\t")
			a++
			goto LOOP
		}


	fmt.Println("asdasdasdasdasd")


}
