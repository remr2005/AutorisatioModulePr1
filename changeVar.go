package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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
		fmt.Println(err)
	}
	defer db.Close()

	//Присваивание значения sets перемнной vars пользователю с айди tel_id
	_, err = db.Exec("update persons.MENSCHEN set "+vars+" = ? where GITID = ?", sets, git_id)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}
	defer db.Close()

	//Присваивание значения sets перемнной vars пользователю с айди tel_id
	_, err = db.Exec("update persons.MENSCHEN set "+vars+" = ? where TELID = ?", sets, tel_id)
	if err != nil {
		fmt.Println(err)
	}

	dec_token, err := jwt.Parse(state, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		fmt.Println(err)
	}
	payload, ok := dec_token.Claims.(jwt.MapClaims)

	if ok && payload["admin"].(bool) && (payload["expires_at"].(float64) >= float64(time.Now().Unix())) {
		tokeExpiresAt := time.Now().Add(time.Minute * time.Duration(15)) // время жизни токена
		payload_second := jwt.MapClaims{
			"id":         payload["id"].(string),
			"admin":      true,
			"expires_at": tokeExpiresAt.Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload_second) // создание нового токена
		NewTokenString, err := token.SignedString([]byte(SECRET))
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "http://localhost:8000/delete?tokenDel="+state+"&tokenSet="+NewTokenString, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "http://localhost:8000/admins/forbidden"+state, http.StatusSeeOther)
	}

	//http.Redirect(w, r, "http://localhost:8000/admins?token="+state, http.StatusSeeOther)
}
