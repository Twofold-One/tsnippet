package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// addr is a flag with the name "addr", default value of ":4000"
	// which value will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// infoLog is a logger for writing information messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog is a logger for writing error massages.
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ldate|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on: %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
