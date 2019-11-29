package bookstoredb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	psqlBookstoreUsername = "psql_bookstore_username"
	psqlBookstorePassword = "psql_bookstore_password"
	psqlBookstoreHost     = "psql_bookstore_host"
	psqlBookstorePort     = "psql_bookstore_port"
	// psql_bookstore_schema   = "psql_bookstore_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(psqlBookstoreUsername)
	password = os.Getenv(psqlBookstorePassword)
	host     = os.Getenv(psqlBookstoreHost)
	port     = os.Getenv(psqlBookstorePort)
	// schema   = os.Getenv(psql_bookstore_schema)
)

const connStrPattern = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

func init() {
	port, err := strconv.ParseUint(port, 10, 32)

	if err != nil {
		fmt.Println("The provided port is not a positive number")
		panic(err)
	}

	connStr := fmt.Sprintf(connStrPattern, host, port, username, password, "bookstore", "disable")
	log.Println("connection to bookstore db ...")

	Client, err = sql.Open("postgres", connStr)

	if err != nil {
		fmt.Printf("err %v", err.Error())
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		fmt.Printf("err %v", err.Error())
		panic(err)
	}

	_, err = Client.Exec(`CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
   			first_name VARCHAR (100),
   			last_name VARCHAR (100),
   			password VARCHAR (100) NOT NULL,
			email VARCHAR (355) UNIQUE NOT NULL,
			updated_on bigint NOT NULL,
			deleted BOOL DEFAULT false NOT NULL,
			status VARCHAR(45) NOT NULL, 
   			created_on bigint NOT NULL)`)

	if err != nil {
		fmt.Printf("err %v", err.Error())
		panic(err)
	}

	log.Println("database successfully configured")
}
