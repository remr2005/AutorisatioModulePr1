package main

import (
	"database/sql"
	"net/http"
)

func changeVars_git(w http.ResponseWriter, r *http.Request) {
	// Git айди
	git_id := r.URL.Query().Get("git_id")
	//Переменная которую будут менять
	vars := r.URL.Query().Get("vars")
	//Значение, которое нужно присвоить перемненной
	sets := r.URL.Query().Get("sets")

	//Открытие БД
	db, err := sql.Open("mysql", "root:godzila2005;@/persons")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Присваивание значения sets перемнной vars пользователю с айди tel_id
	_, err = db.Exec("update persons.MENSCHEN set "+vars+" = ? where GITID = ?", sets, git_id)
	if err != nil {
		panic(err)
	}
}

func changeVars_tel(w http.ResponseWriter, r *http.Request) {
	// Телеграм айди
	tel_id := r.URL.Query().Get("tel_id")
	//Переменная которую будут менять
	vars := r.URL.Query().Get("vars")
	//Значение, которое нужно присвоить перемненной
	sets := r.URL.Query().Get("sets")
	//Значение, которое нужно присвоить перемненной
	state := r.URL.Query().Get("state")

	//Открытие БД
	db, err := sql.Open("mysql", "root:godzila2005;@/persons")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Присваивание значения sets перемнной vars пользователю с айди tel_id
	_, err = db.Exec("update persons.MENSCHEN set "+vars+" = ? where TELID = ?", sets, tel_id)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "http://localhost:8000/admins?token="+state, http.StatusSeeOther)
}
