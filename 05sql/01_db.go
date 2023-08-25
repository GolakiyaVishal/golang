package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         uint16 `json:"id"`
	Name       string `json:"name"`
	Age        uint16 `json:"age"`
	Profession string `json:"profession"`
}

func main() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_db")
	onError("Error in open database", err)
	if err != nil {
		fmt.Printf("Error in open database: %s", err)
	}
	defer db.Close()

	fmt.Println("connected to database")

	res, err := db.Query("SELECT * FROM user")
	onError("Error in select", err)
	defer res.Close()

	for res.Next() {
		var user User

		err := res.Scan(&user.Id, &user.Name, &user.Age, &user.Profession)
		onError("Error on scan", err)

		fmt.Println(user.Name)
	}

	inRes, err := db.Query("INSERT INTO user (name, age, profession) VALUES('Rakesh', 25, 'Software')")
	onError("Error in insert", err)

	defer inRes.Close()

	fmt.Printf("inserted row: %s", "done")
}

func onError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err.Error())
		panic(err.Error())
	}
}
