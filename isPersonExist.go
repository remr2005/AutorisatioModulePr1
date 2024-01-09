package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func isPersonExist_Git(w http.ResponseWriter, r *http.Request) {
	// Принятие параметра
	id := r.URL.Query().Get("id")
	var f bool = false

	db, err := sql.Open("mysql", "root:godzila2005;@/persons") // открытие БД

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() // освобождение памяти

	rows, err := db.Query("select * from persons.MENSCHEN") // получение всех пользователей                               // массив пользователей

	for rows.Next() {
		p := person{}
		err := rows.Scan(&p.GITID, &p.TELID, &p.SURNAME, &p.NAMEP, &p.FATHER_NAME, &p.GROUPP, &p.STUDENT, &p.LEHRER, &p.ADMINP) // запись пользователя
		if err != nil {
			fmt.Println(err)
			continue
		}
		if p.GITID == id {
			f = true
			break
		}
	}

	if f {
		fmt.Fprintf(w, `1`)
	} else {
		fmt.Fprintf(w, `0`)
	}
}

func isPersonExist_Tel(w http.ResponseWriter, r *http.Request) {
	// Принятие параметра
	id := r.URL.Query().Get("id")
	var f bool = false

	db, err := sql.Open("mysql", "root:godzila2005;@/persons") // открытие БД

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() // освобождение памяти

	rows, err := db.Query("select * from persons.MENSCHEN") // получение всех пользователей                               // массив пользователей

	for rows.Next() {
		p := person{}
		err := rows.Scan(&p.GITID, &p.TELID, &p.SURNAME, &p.NAMEP, &p.FATHER_NAME, &p.GROUPP, &p.STUDENT, &p.LEHRER, &p.ADMINP) // запись пользователя
		if err != nil {
			fmt.Println(err)
			continue
		}
		if p.TELID == id {
			f = true
			break
		}
	}

	if f {
		fmt.Fprintf(w, `1`)
	} else {
		fmt.Fprintf(w, `0`)
	}
}
