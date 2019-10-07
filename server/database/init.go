package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// InitRepository creates the database used for the canban board
func InitRepository() {
	db, err := sql.Open("sqlite3", "../server/database/repository/repo.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM todotasks")

	checkErr(err)

	var title, desc, id string

	for rows.Next() {

		err = rows.Scan(&title, &desc, &id)
		checkErr(err)
		fmt.Println(title)
		fmt.Println(desc)
		fmt.Println(id)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
