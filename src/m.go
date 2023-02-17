package main
import "fmt"



func main() {
 
// var variable1 int = 5
// p  := &variable1
// fmt.Println(variable1,p)

// i := 42
// j := &i
// fmt.Println(*j) // читаем значение переменной i через указатель p
// *j = 21         // записываем в переменную i значение 21 через указатель p


incrementCopy := func(i int) {
	i++
}

increment := func(i *int) {
	(*i)++
}

i := 42

incrementCopy(i)
fmt.Println(i) // 42

increment(&i)
fmt.Println(i) // 43




	a := 1
	p := &a
	b := &p
	
	*p = 3
	**b = 5
	
	a += 4 + *p + **b
	
	fmt.Printf("%d",*p)
  

}



