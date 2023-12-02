package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

/* package main

import (
"github.com/gorilla/mux"
"log"
"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/isPersonDelete/git", deletePerson_git).Methods("GET")
	r.HandleFunc("/isPersonDelete/tel", deletePerson_tel).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}
*/

func deletePerson_git(w http.ResponseWriter, r *http.Request) {
	// Github айди
	git_id := r.URL.Query().Get("git_id")

	// Открытие БД
	db, err := sql.Open("mysql", "root:godzila2005;@/persons")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Выполнение SQL-запроса для удаления пользователя по GITID
	result, err := db.Exec("delete from persons.MENSCHEN where GITID = ?", git_id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}

func deletePerson_tel(w http.ResponseWriter, r *http.Request) {
	// Телеграм айди
	tel_id := r.URL.Query().Get("tel_id")

	// Открытие БД
	db, err := sql.Open("mysql", "root:godzila2005;@/persons")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Выполнение SQL-запроса для удаления пользователя по TELID
	result, err := db.Exec("DELETE FROM persons.MENSCHEN WHERE TELID = ?", tel_id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}
