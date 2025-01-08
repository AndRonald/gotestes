package main

import "fmt"

const MAX int = 100
func main(){
	fmt.Println("1 - Oi, estou em uma execução 100% linear")
	fmt.Println("2 - Oi, estou em uma execução 100% linear")
	fmt.Println("3 - Oi, estou em uma execução 100% linear")

	var n int
	n = 10
	
	for {
		n = n + 1;
		fmt.Println(n)
		if n > MAX{
			break;
		}
	}

	switch n < MAX {
		case true: {
			fmt.Println("n é maior que MAX")
		}
		case false: {
			fmt.Println("n não é menor do que max")
		}
	}
	
}

//for ; n < MAX; n = n + 1{
//	fmt.Printf("%d\n", n)
//}