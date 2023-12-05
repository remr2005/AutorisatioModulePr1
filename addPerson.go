package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Person структура представляет пользователя
type Person struct {
	SURNAME     string `json:"SURNAME"`
	NAMEP       string `json:"NAMEP"`
	FATHER_NAME string `json:"FATHER_NAME"`
	GITID       string `json:"GITID"`
	TELID       string `json:"TELID"`
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	// Парсинг JSON-данных из тела запроса в структуру Person
	var newPerson Person
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
	result, err := db.Exec("INSERT INTO persons.MENSCHEN (SURNAME, NAMEP, FATHER_NAME, GITID, TELID) VALUES (?, ?, ?, ?)",
		newPerson.SURNAME, newPerson.NAMEP, newPerson.FATHER_NAME, newPerson.GITID, newPerson.TELID)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId()) // ID последнего добавленного объекта
	fmt.Println(result.RowsAffected()) // Количество затронутых строк
}
