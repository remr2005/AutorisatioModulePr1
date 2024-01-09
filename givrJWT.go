package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func giveJWT(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	tokeExpiresAt := time.Now().Add(time.Minute * time.Duration(3))
	db, err := sql.Open("mysql", "root:godzila2005;@/persons")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Выполнение запроса для получения значения переменной vars пользователя с указанным git_id
	var dbVars string
	err = db.QueryRow("SELECT ADMINP FROM persons.MENSCHEN WHERE GITID = ?", id).Scan(&dbVars)
	if err != nil {
		fmt.Println(err)
	}

	// Проверка совпадения с ролью
	var result bool
	if dbVars == "1" {
		result = true
	} else {
		result = false
	}

	payload := jwt.MapClaims{
		"admin":      result,
		"expires_at": tokeExpiresAt.Unix(), // Не обязательно
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, tokenString)

}
