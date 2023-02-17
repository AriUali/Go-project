package main
import "fmt"

func main() {
 
	var lastWeekTemp [7]int
	tempOnWednesday := lastWeekTemp[3]
	fmt.Println(tempOnWednesday)

// 	thisWeekTemp := [7]int {-3,5,7} // [-3 5 7 0 0 0 0]
//     rgbColor := [3]uint8 {255, 255, 128} // [255 255 128]


// 	rgbColor := [...]uint8{255, 255, 128} // [255 255 128] len = 3
// rgbaColor := [...]uint8{255, 255, 128, 1} // [255 255 128 1] len = 4

// thisWeekTemp := [7]int {6:11, 2:3} // [0 0 3 0 0 0 11]

// var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5} 

// sumTemp := 0

// for i:= 0; i < len(weekTemp); i++ {
//     sumTemp += weekTemp[i]
// }

// average := sumTemp / len(weekTemp)

// fmt.Println(average)


var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5} 

sumTemp := 0

for _, temp := range weekTemp {
    sumTemp += temp
}

average := sumTemp / len(weekTemp)
fmt.Println(average)



weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
workDaysSlice := weekTempArr[:5]
weekendSlice := weekTempArr[5:]
fromTuesdayToThursDaySlice := weekTempArr[1:4] 
weekTempSlice := weekTempArr[:]

fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice)) // [1 2 3 4 5] 5 7
fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice)) // [6 7] 2 2 
fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6 
fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice)) // [1 2 3 4 5 6 7] 7 7

s := []int{1,2,3} // [1 2 3]
s = s[:len(s)-1] // [1 2]

a := []int{1, 2, 3, 4}
b := a[2:3]   // b = [3]
b = append(b, 7)
fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
fmt.Println(b, len(b), cap(b)) // [3 7] 2 2
b = append(b, 8, 9, 10)
b[0] = 11
fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
fmt.Println(b, len(b), cap(b)) // [11 7 8 9 10] 5 6




sEXAMPLE := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7
// 1. Создаём слайс s с базовым массивом на 7 элементов. 
// Четыре первых элемента будут доступны в слайсе.

slice1 := append(sEXAMPLE[:2], 2, 3, 4)  
fmt.Println(sEXAMPLE, slice1) // [0 0 2 3] [0 0 2 3 4]
// 2. Берём слайс из первых двух элементов s и добавляем к ним три элемента.
// Так как суммарная длина полученного слайса (len == 5) меньше ёмкости s[:2] (cap == 7), 
// то базовый массив остаётся прежним.
// Слайс s тоже изменился, но его длина осталась прежней

slice2 := append(sEXAMPLE[1:2], 7) 
fmt.Println(sEXAMPLE, slice1, slice2) // [0 0 7 3] [0 0 7 3 4] [0 7]
// 3. Здесь также базовый массив остаётся прежним, изменились все три слайса

slice3 := append(sEXAMPLE, slice1[1:]...)
fmt.Println(len(slice3), cap(slice3))  // 8 14
// 4. Длина s и slice1[1:] равна 4, длина нового слайса будет равна 8,  
// что больше ёмкости базового массива.
// Будет создан новый базовый массив ёмкостью 14,
// ёмкость нового базового массива подбирается автоматически 
// и зависит от текущего размера и количества добавленных элементов

// 5. Легко проверить, что slice3 ссылается на новый базовый массив
sEXAMPLE[1] = 99
fmt.Println(s, slice1, slice2, slice3) 
// [0 99 7 3] [0 99 7 3 4] [99 7] [0 0 7 3 0 7 3 4]







var dest []int
dest2, dest3 := make([]int, 3),  make([]int, 5)
src := []int{1, 2, 3, 4}
copy(dest, src)
copy(dest2, src)
copy(dest3, src)
fmt.Println(dest, dest2, dest3, src ) // [] [1 2 3] [1 2 3 4 0] [1 2 3 4]



sLS1 := []int{1,2,3,4,5}
i := 2

if len(sLS1) != 0 && i < len(sLS1) { // защищаемся от паники
	s = append(sLS1[:i], sLS1[i+1:]...)
} 
fmt.Println(sLS1) // [1 2 4 5]
}