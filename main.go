package main

import (
	"database/sql"
	"log"
	"migrationtest/api"
	"migrationtest/db"
	"migrationtest/db/userstore"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := sql.Open("postgres", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Fatal("error in connect to db")
	}

	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		log.Fatal("error in ping ")
	}
	defer dbConn.Close()

	newDataBase := db.NewDataBase(dbConn)
	userStore := userstore.NewUserStore(*newDataBase)
	userhandler := api.NewUserHandler(userStore)

	log.Println("server started in localhost:5000 !!🏃‍♂️💨")

	http.HandleFunc("/", userhandler.HandleIndex)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
