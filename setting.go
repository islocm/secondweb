package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgBase     string
	Data       string
	Assets     string
	HTML       string
}

var cfg setting

func init() {
	file, e := os.Open("setting.cfg")
	if e != nil {
		fmt.Println(e.Error())
		panic("ne udalos otkrit fayl konfiguratsiya")
	}
	defer file.Close()
	stat, e := file.Stat()
	if e != nil {
		fmt.Println(e.Error())
		panic("ne udalost prochitat informatsiya o fayle konfigurator")
	}
	readByte := make([]byte, stat.Size())

	_, e = file.Read(readByte)
	if e != nil {
		fmt.Println(e.Error())
		panic("ne udalost fayl prochitat")
	}
	e = json.Unmarshal(readByte, &cfg)
	if e != nil {
		fmt.Println(e.Error())

		panic("ne udalost schitat daniiy fayl konfiguratsiya")
	}
}
