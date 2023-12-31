package main

type person struct {
	GITID       string `json: "GITID"`
	TELID       string `json: "TELID"`
	SURNAME     string `json: "SURNAME"`
	NAMEP       string `json: "NAMEP"`
	FATHER_NAME string `json: "FATHER_NAME"`
	GROUPP      string `json: "GROUPP"`
	STUDENT     uint8  `json: "STUDENT"`
	LEHRER      uint8  `json: "LEHRER"`
	ADMINP      uint8  `json: "ADMINP"`
}

type bufers_struct struct {
	Access_token string `json: "access_token"`
	Token_type   string `json: "token_type"`
	Scope        string `json: "scope"`
}

type git_id struct {
	Id int `json: "id"`
}
