package abc

import (
	"database/sql"
)

func main() {
	db, err := sql.Open("mysql", "sample:sample@tcp(127.0.0.1:3306)/sample")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
