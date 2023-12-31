package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func addPerson(w http.ResponseWriter, r *http.Request) {
	// Парсинг JSON-данных из тела запроса в структуру Person
	var newPerson person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newPerson); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Открытие БД
	db, err := sql.Open("mysql", "root:godzila2005;@/persons")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Выполнение SQL-запроса для добавления нового пользователя
	_, err = db.Exec("INSERT INTO persons.MENSCHEN (GITID, TELID, SURNAME, NAMEP, FATHER_NAME, GROUPP, STUDENT, LEHRER, ADMINP) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		newPerson.GITID, newPerson.TELID, newPerson.SURNAME, newPerson.NAMEP, newPerson.FATHER_NAME, newPerson.GROUPP, newPerson.STUDENT, newPerson.LEHRER, newPerson.ADMINP)
	if err != nil {
		panic(err)
	}
}
