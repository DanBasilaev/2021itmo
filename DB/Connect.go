package DB

import (
	"database/sql"
	"fmt"
)

const(
	host = "localhost"
	port = 5432
	password = "Nad18baS"
	user = "postgres"
	db_name = "Users"
)

func Connect() *sql.DB {
	con, err := sql.Open(user, fmt.Sprintf("host=%s port=%d user=%s "+"password=%s db_name=%s sslmode=disable",
		host, port, user, password, db_name))
	if err != nil{
		panic(err)
	}
	return con
}


