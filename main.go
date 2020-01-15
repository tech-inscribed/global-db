package main

import (
	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/techinscribed/global-db/controllers"
	"github.com/techinscribed/global-db/sqldb"
)

func main() {
	sqldb.ConnectDB()

	http.HandleFunc("/", controllers.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}

	s.ListenAndServe()
}
