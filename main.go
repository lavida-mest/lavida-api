package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"

	_mysql "github.com/muathendirangu/lavida-api/category/mysql"
	_usecase "github.com/muathendirangu/lavida-api/category/usecase"
	"github.com/muathendirangu/lavida-api/server"
)

func main() {

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", "localhost", "3306", "lavida")

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Africa/Accra")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	category := _mysql.New(dbConn)
	categoryUsecase := _usecase.NewService(category)
	srv := server.New(categoryUsecase)

	errs := make(chan error, 2)
	go func() {
		log.Fatalln(http.ListenAndServe(":8080", srv))
		errs <- http.ListenAndServe(":8080", srv)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Fatal("terminated", <-errs)

	fmt.Println("I have linked up ssh is working ")
}
