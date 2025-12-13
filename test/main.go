// go build -ldflags "-s -w"  -- оптимизация сборки
package main

import (
	"fmt"
	//"sort"
	//"math"
)

//самая простая функция
// func sum(a int, b int) int {
// 	result := a + b
// 	return result
// }

//функция с возвратом нескольких переменных
// func math_func(a int, b int) (sum int, subtraction int, mult int, div int) {
// 	sum = a + b
// 	subtraction = a - b
// 	mult = a * b
// 	div = a / b
// 	return
// }

// функция с ссылкой на функцию, находящуюся в главной функции
// func test(some_func func(int) int) {
// 	fmt.Println(some_func(25))
// }

//возвращение функции внутри функции
// func test(x string) func() {
// 	return func() {
// 		fmt.Println(x)
// 	}
// }

// функция для изменения переменной
// func change(str *string) {
// 	*str = "LOL"
// }

// создание структуры
type User struct {
	name     string
	age      int64
	password string
}

func change(u *User) {
	u.name = "Kate"
	u.age = 15
	u.password = "newpass"
}

func main() {
	fmt.Println("Hello, Alexey!")

	//калькулятор
	// fmt.Println("Youve just started a calculator!")
	// fmt.Println("Wat action do u wanna do (+, -, *, /, sqrt, ^) ?")

	// var action string
	// fmt.Scan(&action)

	// var a, b float64
	// fmt.Println("Enter ur 1st number")
	// fmt.Scan(&a)

	// fmt.Println("Enter ur 2nd number")
	// fmt.Scan(&b)

	// switch {
	// case action == "+":
	// 	fmt.Println("a + b = " + fmt.Sprint(a+b))
	// case action == "-":
	// 	fmt.Println("a - b = " + fmt.Sprint(a-b))
	// case action == "*":
	// 	fmt.Println("a * b = " + fmt.Sprint(a*b))
	// case action == "/":
	// 	fmt.Println("a / b = " + fmt.Sprint(a/b))
	// case action == "sqrt":
	// 	fmt.Println("a√b = " + fmt.Sprint(b**&a))
	// case action == "^":
	// 	fmt.Println("a ^ b == " + fmt.Sprint())
	// }

	// введение
	// var (
	// 	age     int    = 18
	// 	message string = "Hello! Im Alexey and Im studin' in MIREA University."
	// )
	// fmt.Println(message)
	// fmt.Println("Im " + fmt.Sprint(age) + " y.o")
	// fmt.Println("нас выебут завтра  ?")
	// var answer string
	// fmt.Scan(&answer)

	// var age1 int
	// fmt.Scan(&age1)
	// if age1 >= 18 {
	// 	fmt.Println("you can buy beer")
	// } else if age1 < 18 {
	// 	fmt.Println("go suck ur paw")
	// }

	//кв ур-е ифами
	// var a, b, c float64
	// fmt.Println("реши кв уравнение")

	// fmt.Println("Введите а")
	// fmt.Scan(&a) //ввод переменной

	// fmt.Println("Введите b")
	// fmt.Scan(&b) //ввод переменной

	// fmt.Println("Введите c")
	// fmt.Scan(&c) //ввод переменной

	// d := (b * b) - 4*(a*c)

	// if d > 0 {
	// 	var x1, x2 float64

	// 	x1 = (-b + math.Sqrt(d)) / (2 * a)
	// 	x2 = (-b - math.Sqrt(d)) / (2 * a)
	// 	fmt.Println("два корня уравнения: ")
	// 	fmt.Println(x1)
	// 	fmt.Println(x2)
	// } else if d == 0 {
	// 	var x float64
	// 	x = (-b) / (2 * a)
	// 	fmt.Println("уравнения имеет один корень")
	// 	fmt.Println(x)
	// } else if d < 0 {
	// 	fmt.Println("уравнение не имеет корней!")
	// }

	// range и цикл
	// nums := []int{1, 2, 3, 4, 5}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	// for index, element := range nums { // цикл range. если не нужна переменная , то написать вместо нее _
	// 	fmt.Printf("Index: %d Element: %d\n", index, element)
	// }

	//матрицы
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// for _, row := range matrix {
	// 	for _, col := range row {
	// 		fmt.Printf("%d ", col)
	// 	}
	// 	fmt.Println()
	// }

	// switch - практичная замена If. табуляция критична!!
	// /w strings
	// name := "Kotlin"

	// switch name {

	// case "Jordan":
	// 	fmt.Println("Hi Jordan")

	// case "Kate":
	// 	fmt.Println("Hi Kate")

	// case "John":
	// 	fmt.Println("Hi John")

	// default:
	// 	fmt.Println("Idk u")
	// }

	// /w nums
	// number := 10

	// switch {
	// case number > 1:
	// 	fmt.Println("Num is > 1")
	// 	fallthrough
	// case number < 9:
	// 	fmt.Println("Num is < 9")
	// }

	//math.Ceil(a) - округление в большую сторону
	//math.Floor(a) - округление в меньшую сторону
	//math.Round(a) - округление по правилам математики

	//массивы  (многомерные массивы дефолт синтаксис)
	// names := [3]string{"kate", "alex", "john"} // увеличить размер массива нельзя

	// for i := 0; i < len(names); i++ { // вывод массива поэлементно
	// 	fmt.Println(names[i])
	// }

	// // ср арифм
	// marks := [5]float64{5, 2, 4, 5, 3}
	// var sum float64 = 0

	// for i := 0; i < len(marks); i++ {
	// 	sum += marks[i]
	// }

	// var res float64 = sum / float64(len(marks))
	// fmt.Println(math.Round(res)) // округлили по правилам математики

	// срезы (динамический массив)
	// slice := []int{3, 1, 2, 7, 4}

	// slice = append(slice, 0) //добавление элемента в срез
	// sort.Ints(slice) // сортировка среза(sort.!!! - !!! тип переменных среза)

	// slice1 := []string{"v", "a", "c", "b", "e"}
	// sort.Strings(slice1)
	// fmt.Println(slice1, slice)

	//наглядное применение range для массива
	// slice := []int{1, 3, 65, 8, 5, 4, 867}
	// for _, element := range slice { //поэлементный вывод среза с помощью функции for-range
	// 	fmt.Printf("%d\n", element)
	// }

	// %d - для форматирования чисел (decimal), %s - для форматирования строк (string), %f - для форматирования float, %t - для форматирования boolean

	// карты
	// var money map[string]int = map[string]int{ // !!!карты сортируют по алфавиту
	// money := map[string]int{
	// 	"dollars": 1000,
	// 	"euros":   2000,
	// 	"apples":  3,
	// }
	// //fmt.Println(money)
	// fmt.Println(money["dollars"])

	// money["dollars"] = 5000 //замена переменной в карте
	// delete(money, "apples") //удаление переменной в карте
	// fmt.Println(money)

	// el, status := money["dollars"] //el - елемент карты по индетификатору, status - переменная boolean (true - if el is exist, false - if its not)
	// fmt.Println(el, status)

	// функции
	// sum(5, 9) // вызов функции
	// value := sum(5, 9)
	// fmt.Println(value)

	// s, su, m, d := math_func(3, 5) // вызов функции с возвратом несокльких переменных
	// fmt.Println(s, su, m, d)

	// a := func(x int, y int) int { // создание функции внутри главной
	// 	return x + y
	// }
	// sum := a(3, 9)
	// fmt.Println(sum)

	//задание функции внутри главной
	// square := func(a int) int {
	// 	return a * a
	// }
	// test(square) // ссылка на функцию

	//test("hello")() //более простой вызов функции

	//указатели
	// a := 5
	// pointer := &a
	// fmt.Println(pointer)  // адрес переменной
	// fmt.Println(*pointer) // значение переменной по адресу
	// s := "LLL"
	// fmt.Println(s)

	//другая запись
	// var point *string = &s
	// change(point)
	// change(&s) // изменяем переменную(без указателя мы просто копируем переменную)
	// fmt.Println(s)
	// num := 9
	// b := &num
	// *b = 35
	// fmt.Println(*b)

	//структуры  (кастомный тип данынх)

	//плохой способ
	// var user User = User{name: "John", age: 23, password: "pass"} //тут можно указывать свойства в любом порядке

	//топ способ
	user := User{"John", 23, "pass"} // тут строго как задавали при объявлении
	fmt.Println(user.name)
	user.name = "Alex"
	fmt.Println(user.name)
	change(&user)
	fmt.Println(user)

}
