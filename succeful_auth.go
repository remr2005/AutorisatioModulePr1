package main

import (
	"fmt"
	"net/http"
)

func succeful_auth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `Успешная регистрация`)
}
