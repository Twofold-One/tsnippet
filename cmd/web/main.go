package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Twofold-One/tsnippet/internal/models"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	// addr is a flag with the name "addr", default value of ":4000"
	// which value will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")
	// dsn (data source name) is a flag with PostgreSQL DSN string
	dsn := flag.String("dsn", "postgres://user:password@localhost:5432/tsnippet", "PostgreSQL data source name")

	flag.Parse()

	// infoLog is a logger for writing information messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog is a logger for writing error massages.
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ldate|log.Lshortfile)

	// db is connection pool to db using postgres driver
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on: %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// openDB wraps sql.Open() and returns a sql.DB connection pool for given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
