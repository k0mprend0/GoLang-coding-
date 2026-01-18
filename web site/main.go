package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// структура для получения данных из БД
type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

// струкртура данных из БД
type Article struct {
	Id                     uint16 // для html первый символ должен быть в верхнем регистре
	Title, Anons, FullText string
}

// переменные для вывода из БД
var posts = []Article{}
var showPost = Article{}

//Функции страниц

// функция главной страницы
func index(w http.ResponseWriter /* параметрр для вывода на странице */, r *http.Request /* отслеживание передачи данных */) {

	// создание из шаблона
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// подключение к БД
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/web_site") // такой код для работы с программой MAMP
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// выборка данных
	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}

	// переменная для хранения данных из БД
	posts = []Article{}

	// обработка БД
	// Next() - True - если если есть след строка, False - если нет
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)

		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
		//fmt.Println(fmt.Sprintf("Post: %s with id %d", post.Title, post.Id))
	}

	tmpl.ExecuteTemplate(w, "index", posts)
}

// функция страницы создания поста
func create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "create", nil)
}

// функция сохранения поста
func save_article(w http.ResponseWriter, r *http.Request) {
	// переменные данных из БД
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	// условия сохранения данных БД
	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Заполните все поля!")
	} else {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/web_site") // такой код для работы с программой MAMP
		if err != nil {
			panic(err)
		}

		defer db.Close()

		// заполняем БД входными данными
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES('%s', '%s', '%s')", title, anons, full_text))

		if err != nil {
			panic(err)
		}

		defer insert.Close()

		// переадресация пользователя
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

// функция страницы конкретного поста
func show_post(w http.ResponseWriter, r *http.Request) {
	// переменная для сохранения Id
	vars := mux.Vars(r)

	// подключение шаблонов
	tmpl, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	// подключение к БД
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/web_site") // такой код для работы с программой MAMP
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// достаем Id из БД
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	// переменная для хранения данных из БД
	showPost = Article{}

	// Next() - True - если если есть след строка, False - если нет
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)

		if err != nil {
			panic(err)
		}

		// контент страницы
		showPost = post
		//fmt.Println(fmt.Sprintf("Post: %s with id %d", post.Title, post.Id))
	}

	tmpl.ExecuteTemplate(w, "show", showPost)
}

// функция отслеживания URL адресов
func handleRequest() {
	rtr := mux.NewRouter() // отслеживание динамических (сложных) URL адресов

	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET") // отслеживание динамических (сложных) URL адресов

	// главная страница
	http.Handle("/", rtr)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) // подключение статических обращений

	// подключение к серверу
	http.ListenAndServe(":1488", rtr)
}

func main() {
	handleRequest()
}
