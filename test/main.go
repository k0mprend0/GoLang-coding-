// go build -ldflags "-s -w"  -- оптимизация сборки
package main

import (
	"fmt"
)

//САМАЯ ПРОСТАЯ ФУНКЦИЯ
// func sum(a int, b int) int {
// 	result := a + b
// 	return result
// }

//ФУНКЦИЯ С ВОЗВРАТОМ НЕСКОЛЬКИХ ПЕРЕМЕННЫХ
// func math_func(a int, b int) (sum int, subtraction int, mult int, div int) {
// 	sum = a + b
// 	subtraction = a - b
// 	mult = a * b
// 	div = a / b
// 	return
// }

// ФУНКЦИЯ С ССЫЛКОЙ НА ФУНКЦИЮ, НАХОДЯЩУЮСЯ ВНУТРИ ГЛАВНОЙ ФУНКЦИИ
// func test(some_func func(int) int) {
// 	fmt.Println(some_func(25))
// }

//ВОЗВРАЩЕНИЕ ФУНКЦИИ ВНУТРИ ФУНКЦИИ
// func test(x string) func() {
// 	return func() {
// 		fmt.Println(x)
// 	}
// }

// ФУНКЦИЯ ДЛЯ ИЗМЕНЕНИЯ ПЕРЕМЕННОЙ
// func change(str *string) {
// 	*str = "LOL"
// }

// СОЗДАНИЕ СТРУКТУРЫ
// type User struct {
// 	name     string
// 	age      int64
// 	password string
// 	score    []int
// }

// func change(u *User) {
// 	u.name = "Kate"
// 	u.age = 15
// 	u.password = "newpass"
// }

// МЕТОД СТРУКТУРЫ
// func (u User) getName() string {
// 	return u.name
// }

// func (u *User) setName(name1 string) {
// 	u.name = name1
// }

// func (u User) isElder() bool {
// 	a := u.age
// 	isTrue := false

// 	if a >= 18 {
// 		isTrue = true
// 	} else if a < 18 {
// 		isTrue = false
// 	}

// 	return isTrue
// }

// func (u User) getHigherScore() int {
// 	higher := 0
// 	for _, sc := range u.score { //первая переменная в range овтечает за нумерацию, второая за элементы
// 		if higher < sc {
// 			higher = sc
// 		}
// 	}
// 	return higher
// }

// ИНТЕРФЕЙСЫ
// type Numbers struct { //создали структуру
// 	num1 int
// 	num2 int
// }

// СОЗДАЛИ ИНТЕРФЕЙС
// type NumbersInterface interface {
// 	Sum() int
// 	Mult() int
// 	Div() float64
// 	Sub() int
// }

// func (n Numbers) Sum() int { //структура унаследовала метод Sum() из интерфейса
// 	return n.num1 + n.num2
// }
// func (n Numbers) Mult() int {
// 	return n.num1 * n.num2
// }
// func (n Numbers) Div() float64 {
// 	return float64(n.num1) / float64(n.num2)
// }
// func (n Numbers) Sub() int {
// 	return n.num1 - n.num2
// }

// ГОРУТИНА (АСИНХРОННЫЕ ФУНКЦИИ)
// func say(greet string, ch chan string, ch2 chan int) {
// 	fmt.Println(greet)
// 	ch <- "Hello from GoRutine" //конструкция для добавления данных в канал
// 	ch2 <- 77
// }

// func say(greet string, ch chan int) {
// 	for i := 0; i <= 5; i++ {
// 		ch <- i
// 	}
// 	close(ch) //закрытие канала
// }

func main() {
	fmt.Println("Hello, Alexey!")

	//КАЛЬКУЛЯТОР
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

	// ВВЕДЕНИЕ
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

	//КВ УР-Е ИФАМИ
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

	// RANGE И ЦИКЛ
	// nums := []int{1, 2, 3, 4, 5}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	//первая переменная в range овтечает за нумерацию, второая за элементы
	// for index, element := range nums { // цикл range. если не нужна переменная , то написать вместо нее _
	// 	fmt.Printf("Index: %d Element: %d\n", index, element)
	// }

	//МАТРИЦЫ
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// for _, row := range matrix {
	// 	for _, col := range row {
	// 		fmt.Printf("%d ", col)
	// 	}
	// 	fmt.Println()
	// }

	// SWITCH - практичная замена If. табуляция критична!!
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

	//МАССИВЫ  (многомерные массивы дефолт синтаксис)
	// names := [3]string{"kate", "alex", "john"} // увеличить размер массива нельзя

	// for i := 0; i < len(names); i++ { // вывод массива поэлементно
	// 	fmt.Println(names[i])
	// }

	// // СР АРИФМ
	// marks := [5]float64{5, 2, 4, 5, 3}
	// var sum float64 = 0

	// for i := 0; i < len(marks); i++ {
	// 	sum += marks[i]
	// }

	// var res float64 = sum / float64(len(marks))
	// fmt.Println(math.Round(res)) // округлили по правилам математики

	// СРЕЗЫ (динамический массив)
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

	// КАРТЫ
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

	// ФУНКЦИИ
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

	//УКАЗАТЕЛИ
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

	//СТРУКТУРЫ  (кастомный тип данынх)

	//плохой способ
	// var user User = User{name: "John", age: 23, password: "pass"} //тут можно указывать свойства в любом порядке

	//топ способ
	// user := User{"John", 23, "pass"} // тут строго как задавали при объявлении
	// fmt.Println(user.name)
	// user.name = "Alex"
	// fmt.Println(user.name)
	// change(&user)
	// fmt.Println(user)

	//ВЫЗОВ СТРУКТУРЫ
	// 	fmt.Println(user.getName())
	// 	user.setName("Michel")
	// 	fmt.Println(user)
	// 	fmt.Println(user.isElder())

	// 	if user.isElder() {
	// 		fmt.Println("come in")
	// 	} else {
	// 		fmt.Println("denied")
	// 	}

	// user := User{"John", 23, "pass", []int{23, 67, 89, 102, 42, 1}}
	// fmt.Println(user.getHigherScore())

	//ИНТЕРФЕЙСЫ
	// var i NumbersInterface   // присвоили интерфейс переменной
	// numbers := Numbers{9, 8} // создали структуру
	// i = numbers
	// fmt.Printf("Сумма: %d\n", i.Sum())
	// fmt.Printf("Произведение: %d\n", i.Mult())
	// fmt.Printf("Частное: %f\n", i.Div())
	// fmt.Printf("Разность: %d\n", i.Sub())

	//РАБОТА С ФАЙЛАМИ
	//создание файла
	//библиотека ioutil (отстой)
	// data := []byte("appending text to new file")
	// e := ioutil.WriteFile("test.txt", data, 0600) // 0600 - дефолтное разрешение

	// if e != nil {
	// 	fmt.Println("Cant create a new file \n", e)
	// }

	//работа с библиотекой os (класс топ)
	// file, err := os.Create("text.txt")
	// if err != nil { //if err exist
	// 	fmt.Println("Cant read a file \n", err)
	// 	os.Exit(1) //стоп программы
	// }
	// defer file.Close() // после проверки на ошибку закрытие соеденения с файлом
	// data := "appending text to a new file"
	// file.WriteString(data)

	// file_data, _ := os.ReadFile(file.Name())
	// fmt.Println(string(file_data))

	// //удаление файла
	// os.Remove("text.txt")

	//СОЗДАНИЕ КАНАЛА
	// ch := make(chan string) // chan -тип данных для канала
	// ch2 := make(chan int)

	// //ВЫЗОВ ГОРУТИНЫ
	// go say("Hello GoLang!", ch, ch2)
	// //time.Sleep(2 * time.Second) //остановка времени основного потока на n ед. времени
	// fmt.Println(<-ch, <-ch2)

	//ЗАКРЫТИЕ КАНАЛА
	// ch := make(chan int)
	// go say("Hello GoLang", ch)
	// for a := range ch {
	// 	fmt.Println(a)
	// }

	//ОПЕРАТОР GOTO (ХРЕНОВЫЙ ОПЕРАТОР/ОПЕРАТОР БЕЗУСЛОВНОГО ПЕРЕХОДА)
	i := 0
LOOP: // создание метки
	if i > 50 {
		goto END
	}

	fmt.Println(i)
	i++
	goto LOOP

END:
}
