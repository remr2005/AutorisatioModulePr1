package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var arr []string = make([]string, 0)

// 5539456894
func auth(w http.ResponseWriter, r *http.Request) {
	tel_id := r.URL.Query().Get("id")
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:45.0) Gecko/20100101 Firefox/45.0")
	if tel_id != "" {
		fmt.Fprint(w, "https://github.com/login/oauth/authorize?client_id=e66f2ecc3c7cdb7af065&redirect_uri=http://localhost:8001/auth&scope=user&response_type=code&state="+tel_id)
		arr = append(arr, tel_id)
	} else if code != "" {
		var f bool = true
		for _, i := range arr {
			if i == state {
				f = false
			}
		}
		if f {
			return
		}

		url := "https://github.com/login/oauth/access_token?client_id=e66f2ecc3c7cdb7af065&client_secret=a931044453b308b713761d8c1b56df99a53d72fa&code=" + code
		method := "POST"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err, 0)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Cookie", "_octo=GH1.1.2027803936.1701809464; logged_in=no")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err, 1)
		}

		var req_ bufers_struct

		err = json.NewDecoder(res.Body).Decode(&req_)
		if err != nil {
			fmt.Println(err, 2)
		}

		url = "https://api.github.com/user"
		method = "GET"

		client = &http.Client{}
		req, err = http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", req_.Token_type+" "+req_.Access_token)
		req.Header.Add("Cookie", "_octo=GH1.1.2027803936.1701809464; logged_in=no")

		res, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		var req_1 git_id
		err = json.NewDecoder(res.Body).Decode(&req_1)

		var per person = person{strconv.Itoa(req_1.Id), state, "", "", "", "", 1, 0, 0}

		db, err := sql.Open("mysql", "root:godzila2005;@/persons")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		// Выполнение SQL-запроса для добавления нового пользователя
		_, err = db.Exec("INSERT INTO persons.MENSCHEN (GITID, TELID, SURNAME, NAMEP, FATHER_NAME, GROUPP, STUDENT, LEHRER, ADMINP) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
			per.GITID, per.TELID, per.SURNAME, per.NAMEP, per.FATHER_NAME, per.GROUPP, per.STUDENT, per.LEHRER, per.ADMINP)
		if err != nil {
			fmt.Println(err)
		}
	}
}
