package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func getMENSCHEN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("1")

	db, err := sql.Open("mysql", "root:godzila2005;@/persons") // открытие БД

	if err != nil {
		panic(err)
	}
	defer db.Close() // освобождение памяти

	rows, err := db.Query("select * from persons.MENSCHEN") // получение всех пользователей
	persons := []person{}                                   // массив пользователей

	for rows.Next() {
		p := person{}
		err := rows.Scan(&p.GITID, &p.TELID, &p.SURNAME, &p.NAMEP, &p.FATHER_NAME, &p.GROUPP, &p.STUDENT, &p.LEHRER, &p.ADMINP) // запись пользователя
		if err != nil {
			fmt.Println(err)
			continue
		}
		persons = append(persons, p) // добавление пользователя в массив
	}
	json.NewEncoder(w).Encode(persons) // возвратить список пользователей на страницу
}

func getMENSCHEN_count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("1")

	i := 0

	db, err := sql.Open("mysql", "root:godzila2005;@/persons") // открытие БД

	if err != nil {
		panic(err)
	}
	defer db.Close() // освобождение памяти

	rows, err := db.Query("select * from persons.MENSCHEN") // получение всех пользователей                                // массив пользователей

	for rows.Next() {
		i+=1
	}
	fmt.Fprint(w,i)
// возвратить список пользователей на страницу
}