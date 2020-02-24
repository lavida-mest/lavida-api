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

	_categoryRepo "github.com/muathendirangu/lavida-api/category/mysql"
	_categoryUsecase "github.com/muathendirangu/lavida-api/category/usecase"
	_guideRepo "github.com/muathendirangu/lavida-api/guide/mysql"
	_guideUsecase "github.com/muathendirangu/lavida-api/guide/usecase"
	"github.com/muathendirangu/lavida-api/server"
	_tripRepo "github.com/muathendirangu/lavida-api/trip/mysql"
	_tripUsecase "github.com/muathendirangu/lavida-api/trip/usecase"
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

	category := _categoryRepo.New(dbConn)
	categoryUsecase := _categoryUsecase.NewService(category)

	guide := _guideRepo.New(dbConn)
	guideUsecase := _guideUsecase.NewService(guide)

	trip := _tripRepo.New(dbConn)
	tripUsecase := _tripUsecase.NewService(trip)

	srv := server.New(categoryUsecase, guideUsecase, tripUsecase)

	errs := make(chan error, 2)
	go func() {
		errs <- http.ListenAndServe(":8080", srv)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Fatal("terminated", <-errs)
}
