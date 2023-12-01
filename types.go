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
