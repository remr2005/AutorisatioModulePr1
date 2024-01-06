package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func checkUserRole_git(w http.ResponseWriter, r *http.Request) {
	// Получение git id из параметров запроса
	git_id := r.URL.Query().Get("git_id")
	// Получение переменной, которую будем проверять
	vars := r.URL.Query().Get("vars")
	// Получение значения, с которым будем сравнивать перемменную
	sets := r.URL.Query().Get("sets")

	// Открытие БД
	db, err := sql.Open("mysql", "root:godzila2005;@/persons")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Выполнение запроса для получения значения переменной vars пользователя с указанным git_id
	var dbVars string
	err = db.QueryRow("SELECT "+vars+" FROM persons.MENSCHEN WHERE GITID = ?", git_id).Scan(&dbVars)
	if err != nil {
		panic(err)
	}

	// Проверка совпадения с ролью
	var result int
	if dbVars == sets {
		result = 1
	} else {
		result = 0
	}

	// Вывод результата
	fmt.Fprintf(w, "%d", result)
}

func checkUserRole_tel(w http.ResponseWriter, r *http.Request) {
	// Получение tel id из параметров запроса
	tel_id := r.URL.Query().Get("tel_id")
	// Получение переменной, которую будем проверять
	vars := r.URL.Query().Get("vars")
	// Получение значения, с которым будем сравнивать перемменную
	sets := r.URL.Query().Get("sets")

	// Открытие БД
	db, err := sql.Open("mysql", "root:godzila2005;@/persons")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Выполнение запроса для получения значения переменной vars пользователя с указанным tel_id
	var dbVars string
	err = db.QueryRow("SELECT "+vars+" FROM persons.MENSCHEN WHERE TELID = ?", tel_id).Scan(&dbVars)
	if err != nil {
		panic(err)
	}

	// Проверка совпадения с ролью
	var result int
	if dbVars == sets {
		result = 1
	} else {
		result = 0
	}

	// Вывод результата
	fmt.Fprintf(w, "%d", result)
}
