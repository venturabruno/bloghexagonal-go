package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/venturabruno/bloghexagonal-go/application/usecase"
	"github.com/venturabruno/bloghexagonal-go/infrastructure/handler"
	"github.com/venturabruno/bloghexagonal-go/infrastructure/middleware"
	"github.com/venturabruno/bloghexagonal-go/infrastructure/persistence"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	dataSorceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dataSorceName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	postRepository := persistence.NewMySQLPostRepository(db)
	postUseCase := usecase.NewPostUseCase(postRepository)

	route := mux.NewRouter()
	negroniHandler := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	handler.MakePostHandlers(route, *negroniHandler, *postUseCase)

	http.Handle("/", route)
	route.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + os.Getenv("WEBSERVER_PORT"),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
