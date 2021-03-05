package main

import (
	"github.com/wailovet/nuwa-sqlited/nuwa_sqlited"
	"log"
)

type TTest struct {
	A string `json:"a"`
}

func main() {
	nuwa_sqlited.SqliteXorm().Config("test.db", "test_", true)
	db := nuwa_sqlited.SqliteXorm().Engine()
	err := db.Sync2(&TTest{})
	log.Println(err)
	_, err = db.Insert(&TTest{
		A: "123",
	})
	log.Println(err)
	tt := TTest{}
	_, err = db.Get(&tt)
	log.Println(err)
	log.Println(tt)
}
