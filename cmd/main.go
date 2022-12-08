package main

import (
	"fmt"
	"privy/config"
	cons "privy/models"
	"privy/routes"
)

func main() {
	config.InitDB()

	echo := routes.GetRoutes()
	addres := cons.Addres
	port := cons.Port

	host := fmt.Sprintf("%s:%s", addres, port)
	_ = echo.Start(host)
}
