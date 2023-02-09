package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		found := false

		// проверяем, что кратно 3, и выводим Fizz
		if i%3 == 0 {
			fmt.Printf("Fizz")
			found = true
		}
		// проверяем, что кратно 5, и выводим Buzz
		if i%5 == 0 {
			fmt.Printf("Buzz")
			found = true
		}
		// если число кратно 3 и 5, выводим FizzBuzz

		if !found {
			// если не выполнилось ни одно из условий, выводим число
			fmt.Println(i)
			continue
		}

		fmt.Println()
	}
}

// func main() {
// var i int
// fmt.Scanln(&i)

// if i > 1946 && i<1964{
// 	fmt.Println("Привет, бумер!")
// } else if i > 1965 && i<1980{
// 	fmt.Println("Привет, представитель X!")

// } else if i > 1981 && i<1996{
// 	fmt.Println("Привет, миллениал!")
// } else if i > 1997 && i<1912 {
// 	fmt.Println("Привет, зумер!")
// } else{
// 	fmt.Println("Привет, my f!")
// }

// num := 0
// for i := 1;i<101; i++{
// 	num++
// 	if num % 3 == 0 && num %5 ==0{
// 		fmt.Println( num , "   ","FIZZbuss")
// 	} else if num % 3 == 0 {
// 		fmt.Println( num , "   ","FIZZ")
// 	}else if  num %5 ==0{
// 		fmt.Println( num , "   ","buss")
// 	} else{
// 		fmt.Println("not")
// 	}
// }

// fmt.Println()

// }
