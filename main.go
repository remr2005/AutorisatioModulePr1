package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/addPerson", addPerson).Methods("POST")
	r.HandleFunc("/personDelete/git", deletePerson_git).Methods("GET")
	r.HandleFunc("/personDelete/tel", deletePerson_tel).Methods("GET")
	r.HandleFunc("/changeVar/git", changeVars_git).Methods("GET")
	r.HandleFunc("/changeVar/tel", changeVars_tel).Methods("GET")
	r.HandleFunc("/getPersons", getMENSCHEN).Methods("GET")
	r.HandleFunc("/isPersonExist/git", isPersonExist_Git).Methods("GET")
	r.HandleFunc("/isPersonExist/tel", isPersonExist_Tel).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}

//r.URL.Query().Get("yourVar")
//r.HandleFunc("/yourURL", yourFunc).Methods("GET")
//func yourFunc(w http.ResponseWriter, r *http.Request)
//fmt.Fprint(w,``)
