package main

import (
	"database/sql"
	"fmt"
	"privy/config"
	"privy/internal/api"
	"privy/internal/repository"
	cons "privy/models"
	"privy/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Dbname
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	repository := repository.New(db)
	handler := api.New(repository)
	echo := routes.GetRoutes(handler)

	addres := cons.Addres
	port := cons.Port
	host := fmt.Sprintf("%s:%s", addres, port)
	_ = echo.Start(host)
}
